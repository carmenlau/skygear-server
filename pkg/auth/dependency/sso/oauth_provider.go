package sso

import (
	"github.com/skygeario/skygear-server/pkg/auth/dependency/loginid"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/urlprefix"
	"github.com/skygeario/skygear-server/pkg/core/config"
	coreTime "github.com/skygeario/skygear-server/pkg/core/time"
)

// OAuthProvider is OAuth 2.0 based provider.
type OAuthProvider interface {
	Type() config.OAuthProviderType
	GetAuthURL(state State, encodedState string) (url string, err error)
	GetAuthInfo(r OAuthAuthorizationResponse, state State) (AuthInfo, error)
}

// NonOpenIDConnectProvider are OAuth 2.0 provider that does not
// implement OpenID Connect or we do not implement yet.
// They are Google, Facebook, Instagram and LinkedIn.
type NonOpenIDConnectProvider interface {
	NonOpenIDConnectGetAuthInfo(r OAuthAuthorizationResponse, state State) (authInfo AuthInfo, err error)
}

// ExternalAccessTokenFlowProvider is provider that the developer
// can somehow acquire an access token and that access token
// can be used to fetch user info.
// They are Google, Facebook, Instagram and LinkedIn.
type ExternalAccessTokenFlowProvider interface {
	ExternalAccessTokenGetAuthInfo(AccessTokenResp) (AuthInfo, error)
}

// OpenIDConnectProvider are OpenID Connect provider.
// They are Azure AD v2.
type OpenIDConnectProvider interface {
	OpenIDConnectGetAuthInfo(r OAuthAuthorizationResponse, state State) (authInfo AuthInfo, err error)
}

type OAuthProviderFactory struct {
	urlPrefixProvider        urlprefix.Provider
	redirectURIFunc          RedirectURLFunc
	tenantConfig             config.TenantConfiguration
	timeProvider             coreTime.Provider
	userInfoDecoder          UserInfoDecoder
	loginIDNormalizerFactory loginid.LoginIDNormalizerFactory
}

func NewOAuthProviderFactory(tenantConfig config.TenantConfiguration, urlPrefixProvider urlprefix.Provider, timeProvider coreTime.Provider, userInfoDecoder UserInfoDecoder, loginIDNormalizerFactory loginid.LoginIDNormalizerFactory, redirectURIFunc RedirectURLFunc) *OAuthProviderFactory {
	return &OAuthProviderFactory{
		tenantConfig:             tenantConfig,
		urlPrefixProvider:        urlPrefixProvider,
		timeProvider:             timeProvider,
		userInfoDecoder:          userInfoDecoder,
		loginIDNormalizerFactory: loginIDNormalizerFactory,
		redirectURIFunc:          redirectURIFunc,
	}
}

func (p *OAuthProviderFactory) NewOAuthProvider(id string) OAuthProvider {
	providerConfig, ok := p.tenantConfig.GetOAuthProviderByID(id)
	if !ok {
		return nil
	}
	switch providerConfig.Type {
	case config.OAuthProviderTypeGoogle:
		return &GoogleImpl{
			URLPrefix:                p.urlPrefixProvider.Value(),
			RedirectURLFunc:          p.redirectURIFunc,
			OAuthConfig:              p.tenantConfig.AppConfig.Identity.OAuth,
			ProviderConfig:           providerConfig,
			TimeProvider:             p.timeProvider,
			UserInfoDecoder:          p.userInfoDecoder,
			LoginIDNormalizerFactory: p.loginIDNormalizerFactory,
		}
	case config.OAuthProviderTypeFacebook:
		return &FacebookImpl{
			URLPrefix:       p.urlPrefixProvider.Value(),
			RedirectURLFunc: p.redirectURIFunc,
			OAuthConfig:     p.tenantConfig.AppConfig.Identity.OAuth,
			ProviderConfig:  providerConfig,
			UserInfoDecoder: p.userInfoDecoder,
		}
	case config.OAuthProviderTypeInstagram:
		return &InstagramImpl{
			URLPrefix:       p.urlPrefixProvider.Value(),
			RedirectURLFunc: p.redirectURIFunc,
			OAuthConfig:     p.tenantConfig.AppConfig.Identity.OAuth,
			ProviderConfig:  providerConfig,
			UserInfoDecoder: p.userInfoDecoder,
		}
	case config.OAuthProviderTypeLinkedIn:
		return &LinkedInImpl{
			URLPrefix:       p.urlPrefixProvider.Value(),
			RedirectURLFunc: p.redirectURIFunc,
			OAuthConfig:     p.tenantConfig.AppConfig.Identity.OAuth,
			ProviderConfig:  providerConfig,
			UserInfoDecoder: p.userInfoDecoder,
		}
	case config.OAuthProviderTypeAzureADv2:
		return &Azureadv2Impl{
			URLPrefix:                p.urlPrefixProvider.Value(),
			RedirectURLFunc:          p.redirectURIFunc,
			OAuthConfig:              p.tenantConfig.AppConfig.Identity.OAuth,
			ProviderConfig:           providerConfig,
			TimeProvider:             p.timeProvider,
			LoginIDNormalizerFactory: p.loginIDNormalizerFactory,
		}
	case config.OAuthProviderTypeApple:
		return &AppleImpl{
			URLPrefix:                p.urlPrefixProvider.Value(),
			RedirectURLFunc:          p.redirectURIFunc,
			OAuthConfig:              p.tenantConfig.AppConfig.Identity.OAuth,
			ProviderConfig:           providerConfig,
			TimeProvider:             p.timeProvider,
			LoginIDNormalizerFactory: p.loginIDNormalizerFactory,
		}
	}
	return nil
}

func (p *OAuthProviderFactory) GetOAuthProviderConfig(id string) (config.OAuthProviderConfiguration, bool) {
	return p.tenantConfig.GetOAuthProviderByID(id)
}
