package server

import (
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/control/router"
	"github.com/gin-gonic/gin"
)

func (srv *Server) BuildRouter() bool {
	gin.DefaultErrorWriter = srv.LoggerWriter
	gin.DefaultWriter = srv.LoggerWriter
	srv.HttpEngine = gin.New()
	srv.HttpEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: srv.LoggerWriter,
	}), gin.Recovery())
	srv.HttpEngine.Use(srv.corsMW)

	srv.Router = router.NewRootRouter(srv.Module)
	srv.Module.Provide(config.ModulePath.Global.HttpEngine, srv.HttpEngine)
	return true
}
