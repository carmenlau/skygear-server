package userverify

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	pkg "github.com/skygeario/skygear-server/pkg/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/authz"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/hook"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal/password"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userprofile"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userverify"
	"github.com/skygeario/skygear-server/pkg/auth/event"
	"github.com/skygeario/skygear-server/pkg/auth/model"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo"
	coreauthz "github.com/skygeario/skygear-server/pkg/core/auth/authz"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/handler"
	"github.com/skygeario/skygear-server/pkg/core/inject"
	"github.com/skygeario/skygear-server/pkg/core/server"
	"github.com/skygeario/skygear-server/pkg/core/validation"
)

// AttachVerifyCodeHandler attaches VerifyCodeHandler to server
func AttachVerifyCodeHandler(
	router *mux.Router,
	authDependency pkg.DependencyMap,
) {
	router.NewRoute().
		Path("/verify_code").
		Handler(server.FactoryToHandler(&VerifyCodeHandlerFactory{
			authDependency,
		})).
		Methods("OPTIONS", "POST")
	router.NewRoute().
		Path("/verify_code_form").
		Handler(server.FactoryToHandler(&VerifyCodeFormHandlerFactory{
			authDependency,
		})).
		Methods("OPTIONS", "POST", "GET")
}

// VerifyCodeHandlerFactory creates VerifyCodeHandler
type VerifyCodeHandlerFactory struct {
	Dependency pkg.DependencyMap
}

// NewHandler creates new VerifyCodeHandler
func (f VerifyCodeHandlerFactory) NewHandler(request *http.Request) http.Handler {
	h := &VerifyCodeHandler{}
	inject.DefaultRequestInject(h, f.Dependency, request)
	return h.RequireAuthz(h, h)
}

type VerifyCodePayload struct {
	Code string `json:"code"`
}

// @JSONSchema
const VerifyCodeRequestSchema = `
{
	"$id": "#VerifyCodeRequest",
	"type": "object",
	"properties": {
		"code": { "type": "string", "minLength": 1 }
	},
	"required": ["code"]
}
`

/*
	@Operation POST /verify_code - Submit verification code
		Verify user using received verification code.

		@Tag User Verification
		@SecurityRequirement access_key
		@SecurityRequirement access_token

		@RequestBody
			@JSONSchema {VerifyCodeRequest}

		@Response 200 {EmptyResponse}

		@Callback user_update {UserUpdateEvent}
		@Callback user_sync {UserSyncEvent}
*/
type VerifyCodeHandler struct {
	TxContext                db.TxContext          `dependency:"TxContext"`
	Validator                *validation.Validator `dependency:"Validator"`
	RequireAuthz             handler.RequireAuthz  `dependency:"RequireAuthz"`
	UserVerificationProvider userverify.Provider   `dependency:"UserVerificationProvider"`
	AuthInfoStore            authinfo.Store        `dependency:"AuthInfoStore"`
	PasswordAuthProvider     password.Provider     `dependency:"PasswordAuthProvider"`
	UserProfileStore         userprofile.Store     `dependency:"UserProfileStore"`
	HookProvider             hook.Provider         `dependency:"HookProvider"`
	Logger                   *logrus.Entry         `dependency:"HandlerLogger"`
}

// ProvideAuthzPolicy provides authorization policy of handler
func (h VerifyCodeHandler) ProvideAuthzPolicy() coreauthz.Policy {
	return authz.AuthAPIRequireValidUser
}

func (h VerifyCodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response handler.APIResponse
	result, err := h.Handle(w, r)
	if err != nil {
		response.Error = err
	} else {
		response.Result = result
	}
	handler.WriteResponse(w, response)
}

func (h VerifyCodeHandler) Handle(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	var payload VerifyCodePayload
	if err := handler.BindJSONBody(r, w, h.Validator, "#VerifyCodeRequest", &payload); err != nil {
		return nil, err
	}

	err = db.WithTx(h.TxContext, func() (err error) {
		userID := auth.GetSession(r.Context()).AuthnAttrs().UserID

		authInfo := &authinfo.AuthInfo{}
		if err = h.AuthInfoStore.GetAuth(userID, authInfo); err != nil {
			return
		}

		var userProfile userprofile.UserProfile
		userProfile, err = h.UserProfileStore.GetUserProfile(userID)
		if err != nil {
			return
		}

		oldUser := model.NewUser(*authInfo, userProfile)

		_, err = h.UserVerificationProvider.VerifyUser(h.PasswordAuthProvider, h.AuthInfoStore, authInfo, payload.Code)
		if err != nil {
			return
		}

		user := model.NewUser(*authInfo, userProfile)

		isVerified := authInfo.IsVerified()
		err = h.HookProvider.DispatchEvent(
			event.UserUpdateEvent{
				Reason:     event.UserUpdateReasonVerification,
				User:       oldUser,
				VerifyInfo: &authInfo.VerifyInfo,
				IsVerified: &isVerified,
			},
			&user,
		)

		resp = struct{}{}
		return
	})
	return
}
