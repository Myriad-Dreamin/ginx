package server

import (
	"github.com/Myriad-Dreamin/ginx/router"
	"github.com/gin-gonic/gin"
)

func (srv *Server) BuildRouter() bool {
	gin.DefaultErrorWriter = srv.LoggerWriter
	gin.DefaultWriter = srv.LoggerWriter
	srv.RouterEngine = gin.New()
	srv.RouterEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: srv.LoggerWriter,
	}), gin.Recovery())
	srv.RouterEngine.Use(srv.corsMW)

	srv.Router = router.NewRootRouter(srv.ServiceProvider, srv.jwtMW, srv.routerAuthMW)
	return true
}
