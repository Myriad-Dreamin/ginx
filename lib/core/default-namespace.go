package mcore

import "path/filepath"

var (
	joiner = filepath.Join
	mini = "minimum"
	global              = joiner(mini, "global")
	middleware          = joiner(mini, "middleware")
	provider            = joiner(mini, "provider")
	dbinstance = joiner(mini, "dbinstance")
	middlewareJWT       = joiner(middleware, "jwt")
	middlewareRouteAuth = joiner(middleware, "route-auth")
	middlewareCORS      = joiner(middleware, "cors")
	globalLogger        = joiner(global, "logger")
	dbinstanceGormDB    = joiner(dbinstance, "gormDB")
	dbinstanceDormDB    = joiner(dbinstance, "dormDB")
	dbinstanceRawDB    = joiner(dbinstance, "rawDB")
	dbinstanceRedisPool = joiner(dbinstance, "redisPool")
	globalConfiguration = joiner(global, "configuration")
	globalHttpEngine    = joiner(global, "httpEngine")
	providerModel       = joiner(provider, "model")
	providerService     = joiner(provider, "service")
	providerRouter      = joiner(provider, "router")
)

var DefaultNamespace = ModulePathS{
	DBInstance: dbInstancesS{
		DormDB: dbinstanceDormDB,
		GormDB: dbinstanceGormDB,
		RawDB: dbinstanceRawDB,
		RedisPool: dbinstanceRedisPool,
	},
	Global: globalS{
		Logger:        globalLogger,
		Configuration: globalConfiguration,
		HttpEngine:    globalHttpEngine,
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

type dbInstancesS struct {
	GormDB string
	DormDB string
	RawDB string
	RedisPool string
}

type globalS struct {
	Logger        string
	Configuration string
	HttpEngine    string
}

type providerS struct {
	Model   string
	Service string
	Router  string
}

type ModulePathS struct {
	Global     globalS
	DBInstance dbInstancesS
	Provider   providerS
	Middleware middlewareS
}
