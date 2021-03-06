package handler

import (
	"github.com/skygeario/skygear-server/pkg/auth/dependency/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/oauth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/oauth/protocol"
	"github.com/skygeario/skygear-server/pkg/core/authn"
	"github.com/skygeario/skygear-server/pkg/core/config"
)

func (h *TokenHandler) IssueAuthAPITokens(
	client config.OAuthClientConfiguration,
	attrs *authn.Attrs,
) (auth.AuthSession, protocol.TokenResponse, error) {
	scopes := []string{"openid", oauth.FullAccessScope}

	authz, err := checkAuthorization(
		h.Authorizations,
		h.Time.NowUTC(),
		h.AppID,
		client.ClientID(),
		attrs.UserID,
		scopes,
	)
	if err != nil {
		return nil, nil, err
	}

	resp := protocol.TokenResponse{}

	offlineGrant, err := h.issueOfflineGrant(client, scopes, authz.ID, attrs, resp)
	if err != nil {
		return nil, nil, err
	}

	err = h.issueAccessGrant(client, scopes, authz.ID,
		offlineGrant.ID, oauth.GrantSessionKindOffline, resp)
	if err != nil {
		return nil, nil, err
	}

	return offlineGrant, resp, nil
}

func (h *TokenHandler) RefreshAPIToken(
	client config.OAuthClientConfiguration,
	refreshToken string,
) (accessToken string, err error) {
	resp, err := h.handleRefreshToken(client, protocol.TokenRequest{
		"client_id":     client.ClientID(),
		"refresh_token": refreshToken,
	})
	if err != nil {
		return "", err
	}
	return resp.GetAccessToken(), nil
}
