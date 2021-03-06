package handler

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"

	pkg "github.com/skygeario/skygear-server/pkg/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/authz"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal"
	"github.com/skygeario/skygear-server/pkg/auth/model"
	coreauthz "github.com/skygeario/skygear-server/pkg/core/auth/authz"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/handler"
	"github.com/skygeario/skygear-server/pkg/core/inject"
	"github.com/skygeario/skygear-server/pkg/core/server"
)

func AttachListIdentitiesHandler(
	router *mux.Router,
	authDependency pkg.DependencyMap,
) {
	router.NewRoute().
		Path("/identity/list").
		Handler(server.FactoryToHandler(&ListIdentitiesHandlerFactory{
			authDependency,
		})).
		Methods("OPTIONS", "POST")
}

type ListIdentitiesHandlerFactory struct {
	Dependency pkg.DependencyMap
}

func (f ListIdentitiesHandlerFactory) NewHandler(request *http.Request) http.Handler {
	h := &ListIdentitiesHandler{}
	inject.DefaultRequestInject(h, f.Dependency, request)
	return h.RequireAuthz(h, h)
}

// @JSONSchema
const IdentityListResponseSchema = `
{
	"$id": "#IdentityListResponse",
	"type": "object",
	"properties": {
		"result": {
			"type": "object",
			"properties": {
				"identities": { 
					"type": "array",
					"items": { "$ref": "#Identity" }
				}
			}
		}
	}
}
`

type IdentityListResponse struct {
	Identities []model.Identity `json:"identities"`
}

/*
	@Operation POST /identity/list - List identities
		Returns list of identities of current user.

		@Tag User
		@SecurityRequirement access_key
		@SecurityRequirement access_token

		@Response 200
			Current user and identity info.
			@JSONSchema {IdentityListResponse}
*/
type ListIdentitiesHandler struct {
	RequireAuthz     handler.RequireAuthz       `dependency:"RequireAuthz"`
	TxContext        db.TxContext               `dependency:"TxContext"`
	IdentityProvider principal.IdentityProvider `dependency:"IdentityProvider"`
}

func (h ListIdentitiesHandler) ProvideAuthzPolicy() coreauthz.Policy {
	return authz.AuthAPIRequireValidUser
}

func (h ListIdentitiesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result, err := h.Handle(w, r)
	if err == nil {
		handler.WriteResponse(w, handler.APIResponse{Result: result})
	} else {
		handler.WriteResponse(w, handler.APIResponse{Error: err})
	}
}

func (h ListIdentitiesHandler) Handle(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	if err = handler.DecodeJSONBody(r, w, &struct{}{}); err != nil {
		return
	}

	err = db.WithTx(h.TxContext, func() error {
		userID := auth.GetSession(r.Context()).AuthnAttrs().UserID

		principals, err := h.IdentityProvider.ListPrincipalsByUserID(userID)
		if err != nil {
			return err
		}

		sort.Slice(principals, func(i, j int) bool {
			return principals[i].PrincipalID() < principals[j].PrincipalID()
		})

		identities := make([]model.Identity, len(principals))
		for i, p := range principals {
			identities[i] = model.NewIdentity(p)
		}

		resp = IdentityListResponse{Identities: identities}
		return nil
	})
	return
}
