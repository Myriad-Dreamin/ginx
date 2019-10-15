package router

import (
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/ginx/service"
)

type RootRouter struct {
	Root       *Router
	Router     *Router
	AuthRouter *Router
	Auth       *Middleware

	ObjectRouter *ObjectRouter
}

func NewRootRouter(serviceProvider *service.Provider, jwtMW *jwt.Middleware, routerAuthMW *Middleware) (r *RootRouter) {
	rr := NewRouterGroup()
	apiRouterV2 := rr.Group("/v1")
	authRouterV2 := apiRouterV2.Group("", jwtMW.Build())

	r = &RootRouter{
		Root:       rr,
		Router:     apiRouterV2,
		AuthRouter: authRouterV2,
		Auth:       routerAuthMW,
	}
	r.ObjectRouter = BuildObjectRouter(r, serviceProvider)
	return
}
