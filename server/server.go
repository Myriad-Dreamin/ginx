package server

import (
	"context"
	"github.com/DeanThompson/ginpprof"
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/lib/plugin"
	"github.com/Myriad-Dreamin/minimum-template/model"
	dblayer "github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-template/router"
	"github.com/Myriad-Dreamin/minimum-template/service"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"io"
	"os"
	"sync"
	"syscall"
)

type Server struct {
	cfg          *config.ServerConfig
	Logger       types.Logger
	LoggerWriter io.Writer

	DB           *gorm.DB
	RedisPool    *redis.Pool
	RouterEngine *gin.Engine
	Router       *router.RootRouter

	contestPath string

	jwtMW *jwt.Middleware
	//var authMW *privileger.MiddleWare
	routerAuthMW *router.Middleware
	corsMW       gin.HandlerFunc

	Module           module.Module
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

type Option interface {
	MinimumServerOption() bool
}

type OptionImpl struct{}

func (OptionImpl) MinimumServerOption() bool { return false }

type OptionRouterLoggerWriter struct {
	OptionImpl
	Writer io.Writer
}

func newServer(options []Option) *Server {
	srv := new(Server)

	for i := range options {
		switch option := options[i].(type) {
		case OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		case *OptionRouterLoggerWriter:
			srv.LoggerWriter = option.Writer
		}
	}

	if srv.LoggerWriter == nil {
		srv.LoggerWriter = os.Stdout
	}

	srv.Module = make(module.Module)
	srv.ServiceProvider = new(service.Provider)
	srv.DatabaseProvider = model.NewProvider("/database")
	srv.RouterProvider = router.NewProvider("/Router")

	_ = model.SetProvider(srv.DatabaseProvider)
	return srv
}

func New(cfgPath string, options ...Option) (srv *Server) {
	srv = newServer(options)
	if !(srv.InstantiateLogger() &&
		srv.LoadConfig(cfgPath) &&
		srv.PrepareDatabase()) {
		srv = nil
		return
	}
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if srv == nil {
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
		srv.println("install router provider error", err)
	}
	if err := srv.Module.Install(srv.DatabaseProvider); err != nil {
		srv.println("install database provider error", err)
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
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if injectSuccess == false {
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
			sugar.PrintStack()
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
