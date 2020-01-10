package router

import (
	"github.com/Myriad-Dreamin/minimum-template/service"
)

type AuthRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	RefreshToken *LeafRouter
}

func BuildAuthRouter(parent *RootRouter, serviceProvider *service.Provider) (router *AuthRouter) {
	authService := serviceProvider.AuthService()

	router = &AuthRouter{
		Router:     parent.GetRouter().Group("auth"),
		AuthRouter: parent.GetAuthRouter().Group("auth"),
		Auth:       parent.GetAuth().Copy(),
	}
	router.RefreshToken = router.AuthRouter.GET("/refresh-token", authService.RefreshToken)

	return
}

func (s *Provider) AuthRouter() *AuthRouter {
	return s.authRouter
}
