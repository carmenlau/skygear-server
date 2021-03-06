package authn

import (
	"github.com/skygeario/skygear-server/pkg/core/skyerr"
)

var (
	InvalidAuthenticationSession  skyerr.Kind = skyerr.Invalid.WithReason("InvalidAuthenticationSession")
	AuthenticationSessionRequired skyerr.Kind = skyerr.Unauthorized.WithReason("AuthenticationSession")
)

var ErrInvalidAuthenticationSession = InvalidAuthenticationSession.New("invalid authentication session")

type oAuthRequireMergeError struct {
	UserID string
}

func (*oAuthRequireMergeError) Error() string { return "require merging oauth user" }
