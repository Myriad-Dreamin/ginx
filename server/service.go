package server

import (
	"github.com/Myriad-Dreamin/ginx/service"
)

func (srv *Server) PrepareService() bool {
	objectService, err := service.NewObjectService(srv.Logger, srv.DatabaseProvider)
	// build Router failed when requesting service with database, report and return
	if err != nil {
		srv.Logger.Debug("get object service error", "error", err)
		return false
	}
	srv.ServiceProvider.Register(objectService)

	return true
}


