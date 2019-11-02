//go:generate package-attach-to-path -generate_register_map
package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
	"github.com/Myriad-Dreamin/ginx/types"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     *model.ObjectDB
	logger types.Logger
}

func NewService(logger types.Logger, provider *model.Provider, _ *config.ServerConfig) (a *Service, err error) {
	a = new(Service)
	a.db = provider.ObjectDB()
	a.logger = logger
	a.CRUDService = base_service.NewCRUDService(a, "oid")
	a.ListService = base_service.NewListService(a, "oid")
	return
}

/*
type Object struct {
}
*/
