package handler

import (
	"context"
	"errors"
	"net/url"
	gotime "time"

	"github.com/sirupsen/logrus"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/oauth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/oauth/protocol"
	"github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/time"
)

const CodeGrantValidDuration = 5 * gotime.Minute

type AuthorizationHandler struct {
	Context context.Context
	AppID   string
	Clients []config.OAuthClientConfiguration
	Logger  *logrus.Entry

	Authorizations       oauth.AuthorizationStore
	CodeGrants           oauth.CodeGrantStore
	AuthorizeEndpoint    oauth.AuthorizeEndpointProvider
	AuthenticateEndpoint oauth.AuthenticateEndpointProvider
	ValidateScopes       ScopesValidator
	CodeGenerator        TokenGenerator
	Time                 time.Provider
}

func (h *AuthorizationHandler) Handle(r protocol.AuthorizationRequest) AuthorizationResult {
	client := resolveClient(h.Clients, r)
	if client == nil {
		return authorizationResultError{
			Response: protocol.NewErrorResponse("unauthorized_client", "invalid client ID"),
		}
	}
	redirectURI, errResp := parseRedirectURI(client, r)
	if errResp != nil {
		return authorizationResultError{Response: errResp}
	}

	result, err := h.doHandle(redirectURI, client, r)
	if err != nil {
		var oauthError *protocol.OAuthProtocolError
		resultErr := authorizationResultError{RedirectURI: redirectURI}
		if errors.As(err, &oauthError) {
			resultErr.Response = oauthError.Response
		} else {
			h.Logger.WithError(err).Error("authz handler failed")
			resultErr.Response = protocol.NewErrorResponse("server_error", "internal server error")
			resultErr.InternalError = true
		}
		state := r.State()
		if state != "" {
			resultErr.Response.State(r.State())
		}
		result = resultErr
	}

	return result
}

func (h *AuthorizationHandler) doHandle(
	redirectURI *url.URL,
	client config.OAuthClientConfiguration,
	r protocol.AuthorizationRequest,
) (AuthorizationResult, error) {
	if err := h.validateRequest(client, r); err != nil {
		return nil, err
	}

	scopes := r.Scope()
	err := h.ValidateScopes(client, scopes)
	if err != nil {
		return nil, err
	}

	session := auth.GetSession(h.Context)
	if session == nil || session.SessionType() != auth.SessionTypeIdentityProvider {
		// Not authenticated as IdP session => request authentication and retry
		return authorizationResultRequireAuthn{
			AuthenticateURI: h.AuthenticateEndpoint.AuthenticateEndpointURI(),
			AuthorizeURI:    h.AuthorizeEndpoint.AuthorizeEndpointURI(),
			Request:         r,
		}, nil
	}

	authz, err := checkAuthorization(
		h.Authorizations,
		h.Time.NowUTC(),
		h.AppID,
		r.ClientID(),
		session.AuthnAttrs().UserID,
		scopes,
	)
	if err != nil {
		return nil, err
	}

	resp := protocol.AuthorizationResponse{}
	switch r.ResponseType() {
	case "code":
		err = h.generateCodeResponse(redirectURI.String(), session, r, authz, scopes, resp)
		if err != nil {
			return nil, err
		}

	case "none":
		break

	default:
		panic("oauth: unexpected response type")
	}

	state := r.State()
	if state != "" {
		resp.State(r.State())
	}

	return authorizationResultRedirect{
		RedirectURI: redirectURI,
		Response:    resp,
	}, nil
}

func (h *AuthorizationHandler) validateRequest(
	client config.OAuthClientConfiguration,
	r protocol.AuthorizationRequest,
) error {
	allowedResponseTypes := client.ResponseTypes()
	if len(allowedResponseTypes) == 0 {
		allowedResponseTypes = []string{"code"}
	}

	ok := false
	for _, respType := range allowedResponseTypes {
		if respType == r.ResponseType() {
			ok = true
			break
		}
	}
	if !ok {
		return protocol.NewError("unauthorized_client", "response type is not allowed for this client")
	}

	if len(r.Scope()) == 0 {
		return protocol.NewError("invalid_request", "scope is required")
	}

	switch r.ResponseType() {
	case "code":
		if r.CodeChallenge() == "" {
			return protocol.NewError("invalid_request", "PKCE code challenge is required")
		}
		if r.CodeChallengeMethod() != "S256" {
			return protocol.NewError("invalid_request", "only 'S256' PKCE transform is supported")
		}
	case "none":
		break
	default:
		return protocol.NewError("unsupported_response_type", "only 'code' response type is supported")
	}

	return nil
}

func (h *AuthorizationHandler) generateCodeResponse(
	redirectURI string,
	session auth.AuthSession,
	r protocol.AuthorizationRequest,
	authz *oauth.Authorization,
	scopes []string,
	resp protocol.AuthorizationResponse,
) error {
	code := h.CodeGenerator()
	codeHash := oauth.HashToken(code)

	codeGrant := &oauth.CodeGrant{
		AppID:           h.AppID,
		AuthorizationID: authz.ID,
		SessionID:       session.SessionID(),

		CreatedAt: h.Time.NowUTC(),
		ExpireAt:  h.Time.NowUTC().Add(CodeGrantValidDuration),
		Scopes:    scopes,
		CodeHash:  codeHash,

		RedirectURI:   redirectURI,
		OIDCNonce:     r.Nonce(),
		PKCEChallenge: r.CodeChallenge(),
	}

	err := h.CodeGrants.CreateCodeGrant(codeGrant)
	if err != nil {
		return err
	}

	resp.Code(code)
	return nil
}
