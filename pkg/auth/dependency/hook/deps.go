package hook

import (
	"context"

	"github.com/google/wire"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal/password"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userprofile"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo"
	"github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/logging"
	"github.com/skygeario/skygear-server/pkg/core/time"
)

func ProvideHookProvider(
	ctx context.Context,
	sqlb db.SQLBuilder,
	sqle db.SQLExecutor,
	requestID logging.RequestID,
	tConfig *config.TenantConfiguration,
	txContext db.TxContext,
	timeProvider time.Provider,
	authInfoStore authinfo.Store,
	userProfileStore userprofile.Store,
	passwordProvider password.Provider,
	loggerFactory logging.Factory,
) Provider {
	return NewProvider(
		ctx,
		string(requestID),
		NewStore(sqlb, sqle),
		txContext,
		timeProvider,
		authInfoStore,
		userProfileStore,
		NewDeliverer(
			tConfig,
			timeProvider,
			NewMutator(
				tConfig.AppConfig.UserVerification,
				passwordProvider,
				authInfoStore,
				userProfileStore,
			),
		),
		loggerFactory,
	)
}

var DependencySet = wire.NewSet(
	ProvideHookProvider,
	wire.Bind(new(auth.HookProvider), new(Provider)),
)
