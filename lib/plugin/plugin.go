package plugin

import (
	"context"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/service"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Logger = types.Logger
type ConfigLoader = types.ConfigLoader
type ServiceProvider = service.Provider
type DatabaseProvider = model.Provider
type ServerConfig = config.ServerConfig
type Module = module.Module

type Plugin interface {
	Configuration(logger Logger, loader ConfigLoader, cfg *ServerConfig) (plg Plugin)
	Inject(services *ServiceProvider, dbs *DatabaseProvider, module Module) (plg Plugin)
	Work(ctx context.Context)
}
