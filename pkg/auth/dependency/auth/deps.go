package auth

import (
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	wire.Struct(new(Middleware), "*"),
	wire.Struct(new(AccessEventProvider), "*"),
	wire.Struct(new(SessionManager), "*"),
)
