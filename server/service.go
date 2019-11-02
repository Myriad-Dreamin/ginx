package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/Myriad-Dreamin/minimum-template/service"
)

type serviceResult struct {
	serviceName string
	functional.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"userService", functional.Decay(service.NewUserService(srv.Logger, srv.DatabaseProvider, srv.jwtMW, srv.cfg))},
		{"objectService", functional.Decay(service.NewObjectService(srv.Logger, srv.DatabaseProvider, srv.cfg))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		srv.ServiceProvider.Register(serviceResult.serviceName, serviceResult.First)
	}
	return true
}
