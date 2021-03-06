package userverify

import (
	"github.com/google/wire"

	"github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/time"
)

var DependencySet = wire.NewSet(
	NewDefaultUserVerifyCodeSenderFactory,
	ProvideProvider,
)

func ProvideProvider(
	tConfig *config.TenantConfiguration,
	time time.Provider,
	builder db.SQLBuilder,
	executor db.SQLExecutor,
) Provider {
	return NewProvider(
		NewCodeGenerator(tConfig),
		NewStore(
			builder,
			executor,
		),
		tConfig.AppConfig.UserVerification,
		time,
	)
}
