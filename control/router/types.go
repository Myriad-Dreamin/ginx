package router

import "github.com/Myriad-Dreamin/minimum-lib/controller"

type Router = controller.Router
type Middleware = controller.Middleware
type LeafRouter = controller.LeafRouter

type H interface {
	GetRouter() *Router
	GetAuthRouter() *Router
	GetAuth() *Middleware
}
