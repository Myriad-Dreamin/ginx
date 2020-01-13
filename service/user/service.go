//go:generate package-attach-to-path -generate_register_map
package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/control"
	base_service "github.com/Myriad-Dreamin/minimum-template/lib/base-service"
	"github.com/Myriad-Dreamin/minimum-template/lib/jwt"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/casbin/casbin/v2"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	userDB     *model.UserDB
	enforcer   *casbin.SyncedEnforcer
	logger     types.Logger
	cfg        *config.ServerConfig
	middleware *jwt.Middleware
}


func (srv *Service) UserServiceSignatureXXX() interface{} { return srv }

func NewService(m module.Module) (a control.UserService, err error) {
	srv := new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	srv.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	srv.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	srv.middleware = m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware)
	srv.userDB = provider.UserDB()
	srv.enforcer = provider.Enforcer()
	srv.CRUDService = base_service.NewCRUDService(srv, srv.cfg.BaseParametersConfig.PathPlaceholder.User)
	srv.ListService = base_service.NewListService(srv, srv.userDB.FilterI)
	a = srv
	return
}

/*
type User struct {
}
*/
