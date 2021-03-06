package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
	gotime "time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/skygeario/skygear-server/pkg/auth/dependency/session"
	coreauth "github.com/skygeario/skygear-server/pkg/core/auth"
	"github.com/skygeario/skygear-server/pkg/core/authn"
	"github.com/skygeario/skygear-server/pkg/core/time"
)

func TestResolveHandler(t *testing.T) {
	Convey("/session/resolve", t, func() {
		h := &ResolveHandler{
			TimeProvider: &time.MockProvider{},
		}

		Convey("should attach headers for valid sessions", func() {
			u := &authn.UserInfo{
				ID:         "user-id",
				IsDisabled: false,
				IsVerified: true,
			}
			d := gotime.Date(2020, 1, 1, 0, 0, 0, 0, gotime.UTC)
			s := &session.IDPSession{
				ID: "session-id",
				Attrs: authn.Attrs{
					PrincipalID:             "principal-id",
					PrincipalType:           "password",
					PrincipalUpdatedAt:      d,
					AuthenticatorID:         "authenticator-id",
					AuthenticatorType:       "oob",
					AuthenticatorOOBChannel: "email",
					AuthenticatorUpdatedAt:  &d,
				},
			}
			r, _ := http.NewRequest("POST", "/", nil)
			r = r.WithContext(authn.WithAuthn(r.Context(), s, u))
			rw := httptest.NewRecorder()
			h.ServeHTTP(rw, r)

			resp := rw.Result()
			So(resp.StatusCode, ShouldEqual, 200)
			So(resp.Header, ShouldResemble, http.Header{
				"X-Skygear-Session-Valid":                     []string{"true"},
				"X-Skygear-User-Id":                           []string{"user-id"},
				"X-Skygear-User-Verified":                     []string{"true"},
				"X-Skygear-User-Disabled":                     []string{"false"},
				"X-Skygear-Session-Identity-Id":               []string{"principal-id"},
				"X-Skygear-Session-Identity-Type":             []string{"password"},
				"X-Skygear-Session-Identity-Updated-At":       []string{"2020-01-01T00:00:00Z"},
				"X-Skygear-Session-Authenticator-Id":          []string{"authenticator-id"},
				"X-Skygear-Session-Authenticator-Type":        []string{"oob"},
				"X-Skygear-Session-Authenticator-Oob-Channel": []string{"email"},
				"X-Skygear-Session-Authenticator-Updated-At":  []string{"2020-01-01T00:00:00Z"},
				"X-Skygear-Is-Master-Key":                     []string{"false"},
			})
		})

		Convey("should attach headers for invalid sessions", func() {
			r, _ := http.NewRequest("POST", "/", nil)
			r = r.WithContext(authn.WithInvalidAuthn(r.Context()))
			rw := httptest.NewRecorder()
			h.ServeHTTP(rw, r)

			resp := rw.Result()
			So(resp.StatusCode, ShouldEqual, 200)
			So(resp.Header, ShouldResemble, http.Header{
				"X-Skygear-Session-Valid": []string{"false"},
				"X-Skygear-Is-Master-Key": []string{"false"},
			})
		})

		Convey("should not attach session headers if no resolved session", func() {
			r, _ := http.NewRequest("POST", "/", nil)
			rw := httptest.NewRecorder()
			h.ServeHTTP(rw, r)

			resp := rw.Result()
			So(resp.StatusCode, ShouldEqual, 200)
			So(resp.Header, ShouldResemble, http.Header{
				"X-Skygear-Is-Master-Key": []string{"false"},
			})
		})

		Convey("should add master key header", func() {
			r, _ := http.NewRequest("POST", "/", nil)
			r = r.WithContext(coreauth.WithAccessKey(r.Context(), coreauth.AccessKey{
				IsMasterKey: true,
			}))
			rw := httptest.NewRecorder()
			h.ServeHTTP(rw, r)

			resp := rw.Result()
			So(resp.StatusCode, ShouldEqual, 200)
			So(resp.Header, ShouldResemble, http.Header{
				"X-Skygear-Is-Master-Key": []string{"true"},
			})
		})
	})
}
