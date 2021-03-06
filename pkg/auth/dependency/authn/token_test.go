package authn

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/skygeario/skygear-server/pkg/core/authn"
)

func TestSessionToken(t *testing.T) {
	Convey("session token", t, func() {
		secret := "secret"
		claims := sessionToken{
			AuthnSession: AuthnSession{
				ClientID:            "clientid",
				RequiredSteps:       []SessionStep{"identity", "mfa"},
				FinishedSteps:       []SessionStep{"identity"},
				SessionCreateReason: "reason",
				Attrs: authn.Attrs{
					UserID:                  "user",
					PrincipalID:             "principal",
					AuthenticatorID:         "authenticator",
					AuthenticatorType:       "totp",
					AuthenticatorOOBChannel: "sms",
				},
			},
		}
		token, err := encodeSessionToken(secret, claims)
		So(err, ShouldBeNil)
		expected, err := decodeSessionToken(secret, token)
		So(err, ShouldBeNil)
		So(&claims, ShouldResemble, expected)
	})
}
