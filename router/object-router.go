package router

import (
	"github.com/Myriad-Dreamin/ginx/service"
)

type ObjectRouter struct {
	*Router
	Auth       *Middleware
	IDRouter   *ObjectIDRouter

	Post    *LeafRouter
	GetList *LeafRouter
}

type ObjectIDRouter struct {
	*Router
	Auth *Middleware

	Get    *LeafRouter
	Put    *LeafRouter
	Delete *LeafRouter
}

func BuildObjectRouter(parent *RootRouter, serviceProvider *service.Provider) (router *ObjectRouter) {
	objectService := serviceProvider.ObjectService()
	router = &ObjectRouter{
		Router:     parent.Router.Extend("object"),
		Auth:       parent.Auth.Copy(),
	}
	router.GetList = router.GET("object-list", objectService.List)
	router.Post = router.POST("/object", objectService.Post)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*ObjectIDRouter) subBuild(parent *ObjectRouter, serviceProvider *service.Provider) (
	router *ObjectIDRouter) {

	objectService := serviceProvider.ObjectService()

	router = &ObjectIDRouter{
		Router: parent.Group("/object/:oid"),
		Auth:   parent.Auth.MustGroup("object", "oid"),
	}

	router.Get = router.GET("", objectService.Get)
	router.Put = router.PUT("", objectService.Put)
	router.Delete = router.DELETE("", objectService.Delete)
	return
}