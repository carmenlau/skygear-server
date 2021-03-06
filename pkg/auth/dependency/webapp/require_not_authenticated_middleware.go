package webapp

import (
	"net/http"

	"github.com/skygeario/skygear-server/pkg/auth/dependency/auth"
)

type RequireNotAuthenticatedMiddleware struct{}

func (m RequireNotAuthenticatedMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := auth.GetUser(r.Context())
		if user != nil {
			RedirectToRedirectURI(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
