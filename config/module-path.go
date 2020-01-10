package config

import "path/filepath"

var (
	global              = "global"
	middleware          = "middleware"
	provider            = "provider"
	middlewareJWT       = filepath.Join(middleware, "jwt")
	middlewareRouteAuth = filepath.Join(middleware, "route-auth")
	middlewareCORS      = filepath.Join(middleware, "cors")
	globalCities        = filepath.Join(global, "cities")
	globalLogger        = filepath.Join(global, "logger")
	globalGormDB        = filepath.Join(global, "gormDB")
	globalConfiguration = filepath.Join(global, "configuration")
	globalHttpEngine    = filepath.Join(global, "httpEngine")
	providerModel       = filepath.Join(provider, "model")
	providerService     = filepath.Join(provider, "service")
	providerRouter      = filepath.Join(provider, "router")
)

var ModulePath = ModulePathS{
	Global: globalS{
		Logger:        globalLogger,
		GormDB:        globalGormDB,
		Configuration: globalConfiguration,
		HttpEngine:    globalHttpEngine,
		Cities:        globalCities,
	},
	Provider: providerS{
		Model:   providerModel,
		Service: providerService,
		Router:  providerRouter,
	},
	Middleware: middlewareS{
		JWT:       middlewareJWT,
		RouteAuth: middlewareRouteAuth,
		CORS:      middlewareCORS,
	},
}

type middlewareS struct {
	JWT       string
	RouteAuth string
	CORS      string
}

type globalS struct {
	Logger        string
	Configuration string
	GormDB        string
	HttpEngine    string
	Cities        string
}

type providerS struct {
	Model   string
	Service string
	Router  string
}

type ModulePathS struct {
	Global     globalS
	Provider   providerS
	Middleware middlewareS
}
