package mfa

import (
	"net/http"

	"github.com/skygeario/skygear-server/pkg/auth/dependency/authn"
	coreauthn "github.com/skygeario/skygear-server/pkg/core/authn"
	"github.com/skygeario/skygear-server/pkg/core/config"
)

type authnResolver interface {
	Resolve(
		client config.OAuthClientConfiguration,
		authnSessionToken string,
		stepPredicate func(authn.SessionStep) bool,
	) (*authn.AuthnSession, error)
}

type authnStepper interface {
	StepSession(
		client config.OAuthClientConfiguration,
		s coreauthn.Attributer,
		mfaBearerToken string,
	) (authn.Result, error)

	WriteAPIResult(rw http.ResponseWriter, result authn.Result)
}
