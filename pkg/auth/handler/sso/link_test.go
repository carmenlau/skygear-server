package sso

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	authtesting "github.com/skygeario/skygear-server/pkg/auth/dependency/auth/testing"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/sso"
	coreconfig "github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/db"
	. "github.com/skygeario/skygear-server/pkg/core/skytest"
	"github.com/skygeario/skygear-server/pkg/core/validation"
)

func TestLinkHandler(t *testing.T) {
	Convey("Test LinkHandler", t, func() {
		stateJWTSecret := "secret"
		providerName := "mock"
		providerUserInfo := sso.ProviderUserInfo{
			ID:    "mock_user_id",
			Email: "john.doe@example.com",
		}

		sh := &LinkHandler{}
		sh.TxContext = db.NewMockTxContext()
		validator := validation.NewValidator("http://v2.skygear.io")
		validator.AddSchemaFragments(
			LinkRequestSchema,
		)
		sh.Validator = validator
		oauthConfig := &coreconfig.OAuthConfiguration{
			StateJWTSecret:                 stateJWTSecret,
			ExternalAccessTokenFlowEnabled: true,
		}
		providerConfig := coreconfig.OAuthProviderConfiguration{
			ID:           providerName,
			Type:         "google",
			ClientID:     "mock_client_id",
			ClientSecret: "mock_client_secret",
		}
		mockProvider := sso.MockSSOProvider{
			URLPrefix:       &url.URL{Scheme: "https", Host: "api.example.com"},
			RedirectURLFunc: RedirectURIForAPI,
			BaseURL:         "http://mock/auth",
			OAuthConfig:     oauthConfig,
			ProviderConfig:  providerConfig,
			UserInfo:        providerUserInfo,
		}
		sh.OAuthProvider = &mockProvider
		sh.SSOProvider = &mockProvider

		Convey("should reject payload without access token", func() {
			req, _ := http.NewRequest("POST", "", strings.NewReader(`{}`))
			req = authtesting.WithAuthn().
				UserID("faseng.cat.id").
				PrincipalID("faseng.cat.principal.id").
				ToRequest(req)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			sh.ServeHTTP(resp, req)
			So(resp.Code, ShouldEqual, 400)
			So(resp.Body.Bytes(), ShouldEqualJSON, `{
				"error": {
					"name": "Invalid",
					"reason": "ValidationFailed",
					"message": "invalid request body",
					"code": 400,
					"info": {
						"causes": [
							{
								"kind": "Required",
								"message": "access_token is required",
								"pointer": "/access_token"
							}
						]
					}
				}
			}`)
		})

		Convey("should return error if disabled", func() {
			mockProvider.OAuthConfig.ExternalAccessTokenFlowEnabled = false
			req, _ := http.NewRequest("POST", "", strings.NewReader(`{
                               "access_token": "token"
                       }`))
			req = authtesting.WithAuthn().
				UserID("faseng.cat.id").
				PrincipalID("faseng.cat.principal.id").
				ToRequest(req)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			sh.ServeHTTP(resp, req)
			So(resp.Code, ShouldEqual, 404)
			So(resp.Body.Bytes(), ShouldEqualJSON, `{
				"error": {
					"name": "NotFound",
					"reason": "NotFound",
					"message": "external access token flow is disabled",
					"code": 404
				}
			}`)
		})
	})
}
