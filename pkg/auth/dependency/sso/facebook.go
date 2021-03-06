package sso

import (
	"net/url"

	"github.com/skygeario/skygear-server/pkg/core/config"
)

const (
	facebookAuthorizationURL string = "https://www.facebook.com/v6.0/dialog/oauth"
	// nolint: gosec
	facebookTokenURL    string = "https://graph.facebook.com/v6.0/oauth/access_token"
	facebookUserInfoURL string = "https://graph.facebook.com/v6.0/me?fields=id,email"
)

type FacebookImpl struct {
	URLPrefix       *url.URL
	RedirectURLFunc RedirectURLFunc
	OAuthConfig     *config.OAuthConfiguration
	ProviderConfig  config.OAuthProviderConfiguration
	UserInfoDecoder UserInfoDecoder
}

func (f *FacebookImpl) Type() config.OAuthProviderType {
	return config.OAuthProviderTypeFacebook
}

func (f *FacebookImpl) GetAuthURL(state State, encodedState string) (string, error) {
	p := authURLParams{
		oauthConfig:    f.OAuthConfig,
		redirectURI:    f.RedirectURLFunc(f.URLPrefix, f.ProviderConfig),
		providerConfig: f.ProviderConfig,
		encodedState:   encodedState,
		baseURL:        facebookAuthorizationURL,
	}
	if state.UXMode == UXModeWebPopup {
		// https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow
		p.display = "popup"
	}
	return authURL(p)
}

func (f *FacebookImpl) GetAuthInfo(r OAuthAuthorizationResponse, state State) (authInfo AuthInfo, err error) {
	return f.NonOpenIDConnectGetAuthInfo(r, state)
}

func (f *FacebookImpl) NonOpenIDConnectGetAuthInfo(r OAuthAuthorizationResponse, state State) (authInfo AuthInfo, err error) {
	h := getAuthInfoRequest{
		redirectURL:     f.RedirectURLFunc(f.URLPrefix, f.ProviderConfig),
		oauthConfig:     f.OAuthConfig,
		providerConfig:  f.ProviderConfig,
		accessTokenURL:  facebookTokenURL,
		userProfileURL:  facebookUserInfoURL,
		userInfoDecoder: f.UserInfoDecoder,
	}
	return h.getAuthInfo(r, state)
}

func (f *FacebookImpl) ExternalAccessTokenGetAuthInfo(accessTokenResp AccessTokenResp) (authInfo AuthInfo, err error) {
	h := getAuthInfoRequest{
		redirectURL:     f.RedirectURLFunc(f.URLPrefix, f.ProviderConfig),
		oauthConfig:     f.OAuthConfig,
		providerConfig:  f.ProviderConfig,
		accessTokenURL:  facebookTokenURL,
		userProfileURL:  facebookUserInfoURL,
		userInfoDecoder: f.UserInfoDecoder,
	}
	return h.getAuthInfoByAccessTokenResp(accessTokenResp)
}

var (
	_ OAuthProvider                   = &FacebookImpl{}
	_ NonOpenIDConnectProvider        = &FacebookImpl{}
	_ ExternalAccessTokenFlowProvider = &FacebookImpl{}
)
