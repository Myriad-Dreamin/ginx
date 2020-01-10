package mgin

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinRoutes struct {
	e gin.IRoutes
}

func (g GinRoutes) Use(handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.Use(ToGinHandlers(handlers)...)}
}

func (g GinRoutes) Handle(method string, path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.Handle(method, path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) Any(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.Any(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) GET(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.GET(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) POST(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.POST(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) DELETE(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.DELETE(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) PATCH(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.PATCH(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) PUT(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.PUT(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) OPTIONS(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.OPTIONS(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) HEAD(path string, handlers ...controller.HandlerFunc) controller.IRoutes {
	return GinRoutes{g.e.HEAD(path, ToGinHandlers(handlers)...)}
}

func (g GinRoutes) StaticFile(relativePath, filepath string) controller.IRoutes {
	return GinRoutes{g.e.StaticFile(relativePath, filepath)}
}

func (g GinRoutes) Static(relativePath, root string) controller.IRoutes {
	return GinRoutes{g.e.Static(relativePath, root)}
}

func (g GinRoutes) StaticFS(relativePath string, fs http.FileSystem) controller.IRoutes {
	return GinRoutes{g.e.StaticFS(relativePath, fs)}
}

type GinRouter struct {
	GinRoutes
	e gin.IRouter
}

func NewGinRouter(e gin.IRouter) GinRouter {
	return GinRouter{GinRoutes{e}, e}
}

func (g GinRouter) Group(path string, handlers ...controller.HandlerFunc) controller.IRouter {
	return NewGinRouter(g.e.Group(path, ToGinHandlers(handlers)...))
}
