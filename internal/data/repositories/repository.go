package repositories

import (
	"demo/internal/data/repositories/impl"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(impl.NewGreeterRepo, impl.NewUserRepository)
