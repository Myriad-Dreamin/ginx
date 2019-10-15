package server

import (
	"github.com/Myriad-Dreamin/ginx/router"
	"github.com/gin-gonic/gin"
)



func (srv *Server) BuildRouter() bool {

	srv.RouterEngine = gin.Default()
	srv.RouterEngine.Use(srv.corsMW)

	srv.RouterEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv.Router = router.NewRootRouter(srv.ServiceProvider, srv.jwtMW, srv.routerAuthMW)
	return true
}
