package router

import (
	"github.com/Myriad-Dreamin/ginx/service"
)

type UserRouter struct {
	*Router
	Auth     *Middleware
	IDRouter *UserIDRouter

	Login    *LeafRouter
	Register *LeafRouter
	GetList  *LeafRouter
}

type UserIDRouter struct {
	*Router
	Auth *Middleware

	ChangePassword *LeafRouter
	Get            *LeafRouter
	Put            *LeafRouter
	Delete         *LeafRouter
}

func BuildUserRouter(parent *RootRouter, serviceProvider *service.Provider) (router *UserRouter) {
	userService := serviceProvider.UserService()
	router = &UserRouter{
		Router: parent.Router.Extend("user"),
		Auth:   parent.Auth.Copy(),
	}
	router.GetList = router.GET("user-list", userService.List)
	router.Register = router.POST("/user", userService.Register)
	router.Login = router.POST("/login", userService.Login)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*UserIDRouter) subBuild(parent *UserRouter, serviceProvider *service.Provider) (
	router *UserIDRouter) {

	userService := serviceProvider.UserService()

	router = &UserIDRouter{
		Router: parent.Group("/user/:id"),
		Auth:   parent.Auth.MustGroup("user", "id"),
	}

	router.Get = router.GET("", userService.Get)
	router.ChangePassword = router.POST("/password", userService.ChangePassword)
	router.Put = router.PUT("", userService.Put)
	router.Delete = router.DELETE("", userService.Delete)
	return
}

func (s *Provider) UserRouter() *UserRouter {
	return s.userRouter
}
