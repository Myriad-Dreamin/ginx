//go:generate package-attach-to-path -generate_register_map
package userservice

import (
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/control"
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/casbin/casbin/v2"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db         *model.UserDB
	enforcer   *casbin.SyncedEnforcer
	logger     types.Logger
	middleware *jwt.Middleware
}

func NewService(logger types.Logger, provider *model.Provider, middleware *jwt.Middleware, _ *config.ServerConfig) (a control.UserService, err error) {
	srv := new(Service)
	srv.db = provider.UserDB()
	srv.enforcer = provider.Enforcer()
	srv.logger = logger
	srv.middleware = middleware
	srv.CRUDService = base_service.NewCRUDService(srv, "id")
	srv.ListService = base_service.NewListService(srv, "id")
	a = srv
	return
}

/*
type User struct {
}
*/
