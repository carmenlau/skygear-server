// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package task

import (
	"context"
	"github.com/google/wire"
	"github.com/skygeario/skygear-server/pkg/auth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/audit"
	pq2 "github.com/skygeario/skygear-server/pkg/auth/dependency/passwordhistory/pq"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal/oauth"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/principal/password"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userprofile"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userverify"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/welcemail"
	"github.com/skygeario/skygear-server/pkg/core/async"
	"github.com/skygeario/skygear-server/pkg/core/auth/authinfo/pq"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/logging"
	"github.com/skygeario/skygear-server/pkg/core/mail"
	"github.com/skygeario/skygear-server/pkg/core/sms"
	"github.com/skygeario/skygear-server/pkg/core/time"
)

// Injectors from wire.go:

func newWelcomeEmailSendTask(ctx context.Context, m auth.DependencyMap) async.Task {
	tenantConfiguration := auth.ProvideTenantConfig(ctx)
	sender := mail.ProvideMailSender(ctx, tenantConfiguration)
	engine := auth.ProvideTemplateEngine(tenantConfiguration, m)
	welcemailSender := welcemail.NewDefaultSender(tenantConfiguration, sender, engine)
	provider := time.NewProvider()
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(ctx, tenantConfiguration)
	store := userprofile.ProvideStore(provider, sqlBuilder, sqlExecutor)
	txContext := db.ProvideTxContext(ctx, tenantConfiguration)
	requestID := ProvideLoggingRequestID(ctx)
	factory := logging.ProvideLoggerFactory(ctx, requestID, tenantConfiguration)
	welcomeEmailSendTask := &WelcomeEmailSendTask{
		WelcomeEmailSender: welcemailSender,
		UserProfileStore:   store,
		TxContext:          txContext,
		LoggerFactory:      factory,
	}
	return welcomeEmailSendTask
}

func newVerifyCodeSendTask(ctx context.Context, m auth.DependencyMap) async.Task {
	tenantConfiguration := auth.ProvideTenantConfig(ctx)
	engine := auth.ProvideTemplateEngine(tenantConfiguration, m)
	sender := mail.ProvideMailSender(ctx, tenantConfiguration)
	client := sms.ProvideSMSClient(ctx, tenantConfiguration)
	codeSenderFactory := userverify.NewDefaultUserVerifyCodeSenderFactory(tenantConfiguration, engine, sender, client)
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlExecutor := db.ProvideSQLExecutor(ctx, tenantConfiguration)
	store := pq.ProvideStore(sqlBuilderFactory, sqlExecutor)
	provider := time.NewProvider()
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	userprofileStore := userprofile.ProvideStore(provider, sqlBuilder, sqlExecutor)
	userverifyProvider := userverify.ProvideProvider(tenantConfiguration, provider, sqlBuilder, sqlExecutor)
	passwordhistoryStore := pq2.ProvidePasswordHistoryStore(provider, sqlBuilder, sqlExecutor)
	requestID := ProvideLoggingRequestID(ctx)
	factory := logging.ProvideLoggerFactory(ctx, requestID, tenantConfiguration)
	reservedNameChecker := auth.ProvideReservedNameChecker(m)
	passwordProvider := password.ProvidePasswordProvider(sqlBuilder, sqlExecutor, provider, passwordhistoryStore, factory, tenantConfiguration, reservedNameChecker)
	oauthProvider := oauth.ProvideOAuthProvider(sqlBuilder, sqlExecutor)
	v := auth.ProvidePrincipalProviders(oauthProvider, passwordProvider)
	identityProvider := principal.ProvideIdentityProvider(sqlBuilder, sqlExecutor, v)
	txContext := db.ProvideTxContext(ctx, tenantConfiguration)
	verifyCodeSendTask := &VerifyCodeSendTask{
		CodeSenderFactory:        codeSenderFactory,
		AuthInfoStore:            store,
		UserProfileStore:         userprofileStore,
		UserVerificationProvider: userverifyProvider,
		PasswordAuthProvider:     passwordProvider,
		IdentityProvider:         identityProvider,
		TxContext:                txContext,
		LoggerFactory:            factory,
	}
	return verifyCodeSendTask
}

func newPwHouseKeeperTask(ctx context.Context, m auth.DependencyMap) async.Task {
	tenantConfiguration := auth.ProvideTenantConfig(ctx)
	txContext := db.ProvideTxContext(ctx, tenantConfiguration)
	requestID := ProvideLoggingRequestID(ctx)
	factory := logging.ProvideLoggerFactory(ctx, requestID, tenantConfiguration)
	provider := time.NewProvider()
	sqlBuilderFactory := db.ProvideSQLBuilderFactory(tenantConfiguration)
	sqlBuilder := auth.ProvideAuthSQLBuilder(sqlBuilderFactory)
	sqlExecutor := db.ProvideSQLExecutor(ctx, tenantConfiguration)
	store := pq2.ProvidePasswordHistoryStore(provider, sqlBuilder, sqlExecutor)
	pwHousekeeper := audit.ProvidePwHousekeeper(store, factory, tenantConfiguration)
	pwHousekeeperTask := &PwHousekeeperTask{
		TxContext:     txContext,
		LoggerFactory: factory,
		PwHousekeeper: pwHousekeeper,
	}
	return pwHousekeeperTask
}

// wire.go:

var DependencySet = wire.NewSet(
	ProvideLoggingRequestID,
)
