package server

import (
	"context"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/model"
	dblayer "github.com/Myriad-Dreamin/ginx/model/db-layer"
	"github.com/Myriad-Dreamin/ginx/plugin"
	"github.com/Myriad-Dreamin/ginx/router"
	"github.com/Myriad-Dreamin/ginx/service"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"runtime"
	"sync"
	"syscall"
)

func printStack() {
	var buf [1024 * 10]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

type Server struct {
	cfg    *config.ServerConfig
	Logger types.Logger

	DB           *gorm.DB
	RedisPool    *redis.Pool
	RouterEngine *gin.Engine
	Router       *router.RootRouter

	contestPath string

	jwtMW *jwt.Middleware
	//var authMW *privileger.MiddleWare
	routerAuthMW *router.Middleware
	corsMW       gin.HandlerFunc

	Module           types.Module
	ServiceProvider  *service.Provider
	DatabaseProvider *model.Provider
	RouterProvider   *router.Provider

	plugins []plugin.Plugin
}

func (srv *Server) Terminate() {
	err := srv.DB.Close()
	if err != nil {
		srv.Logger.Error("close DB error", "error", err)
	}
	syscall.Exit(0)
}

func New(cfgPath string) (srv *Server) {
	srv = new(Server)

	srv.Module = make(types.Module)
	srv.ServiceProvider = new(service.Provider)
	srv.DatabaseProvider = model.NewProvider("/database")
	srv.RouterProvider = router.NewProvider("/Router")

	_ = model.SetProvider(srv.DatabaseProvider)

	if !(srv.InstantiateLogger() &&
		srv.LoadConfig(cfgPath) &&
		srv.PrepareDatabase()) {
		srv = nil
		return
	}
	defer func() {
		if err := recover(); err != nil {
			printStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}

		if srv == nil {
			srv.Terminate()
		}
	}()

	if !(srv.PrepareMiddleware() &&
		srv.PrepareService() &&
		srv.BuildRouter()) {
		srv = nil
		return
	}

	if err := srv.Module.Install(srv.RouterProvider); err != nil {
		fmt.Println("install Router provider error", err)
	}
	if err := srv.Module.Install(srv.DatabaseProvider); err != nil {
		fmt.Println("install database provider error", err)
	}
	//
	//if !PreparePlugin(cfg) {
	//	srv = nil
	//return
	//}

	// Pressure()
	return
}

func (srv *Server) Inject(plugins ...plugin.Plugin) (injectSuccess bool) {
	defer func() {

		if err := recover(); err != nil {
			printStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}

		if injectSuccess == false {
			srv.Terminate()
		}
	}()

	for _, plg := range plugins {
		plg = plg.Configuration(srv.Logger, srv.FetchConfig, srv.cfg)
		if plg == nil {
			return false
		}
		plg = plg.Inject(srv.ServiceProvider, srv.DatabaseProvider, srv.Module)
		if plg == nil {
			return false
		}
		srv.plugins = append(srv.plugins, plg)
	}
	return true
}

func (srv *Server) Serve(port string) {
	defer func() {
		if err := recover(); err != nil {
			printStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}
	}()

	srv.Router.Root.Build(srv.RouterEngine)
	srv.Module.Debug(srv.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	for _, plg := range srv.plugins {
		go plg.Work(ctx)
	}

	if err := dblayer.GetRawInstance().Ping(); err != nil {
		srv.Logger.Debug("database died", "error", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := srv.RouterEngine.Run(port); err != nil {
			srv.Logger.Debug("Router run error", "error", err)
		}
		wg.Done()
	}()

	//do something
	wg.Wait()
}

func (srv *Server) ServeWithPProf(port string) {
	ginpprof.Wrap(srv.RouterEngine)
	srv.Serve(port)
}
