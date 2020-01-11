package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

type dbResult struct {
	dbName string
	functional.DecayResult
}

func (srv *Server) registerDatabaseService() bool {

	for _, dbResult := range []dbResult{
		{"userDB", functional.Decay(model.NewUserDB(srv.Module))},
		{"objectDB", functional.Decay(model.NewObjectDB(srv.Module))},
	} {
		if dbResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("init %T DB error", dbResult.First), "error", dbResult.Err)
			return false
		}
		srv.ModelProvider.Register(dbResult.dbName, dbResult.First)
	}
	return true
}

func (srv *Server) PrepareDatabase() bool {
	var err error
	cfg := srv.Cfg
	srv.DB, err = model.OpenORM(cfg)
	if err != nil {
		srv.Logger.Error("open database error", "error", err)
		return false
	}
	srv.Module.Provide(config.ModulePath.DBInstance.GormDB, srv.DB)

	srv.Logger.Info("connected to database",
		"connection-type", cfg.DatabaseConfig.ConnectionType,
		"user", cfg.DatabaseConfig.User,
		"database", cfg.DatabaseConfig.DatabaseName,
		"charset", cfg.DatabaseConfig.Charset,
		"location", cfg.DatabaseConfig.Location,
	)

	if !model.Register(srv.Module) {
		return false
	}

	//srv.RedisPool, err = model.OpenRedis(cfg)
	//if err != nil {
	//	srv.Logger.Debug("create redis pool error", "error", err)
	//	return false
	//}
	//
	//srv.Logger.Info("connected to redis",
	//	"connection-type", cfg.RedisConfig.ConnectionType,
	//	"host", cfg.RedisConfig.Host,
	//	"connection-timeout", cfg.RedisConfig.ConnectionTimeout,
	//	"database", cfg.RedisConfig.Database,
	//	"read-timeout", cfg.RedisConfig.ReadTimeout,
	//	"write-timeout", cfg.RedisConfig.WriteTimeout,
	//	"idle-timeout", cfg.RedisConfig.IdleTimeout,
	//	"wait", cfg.RedisConfig.Wait,
	//	"max-active", cfg.RedisConfig.MaxActive,
	//	"max-idle", cfg.RedisConfig.MaxIdle,
	//)
	//err = model.RegisterRedis(srv.RedisPool, srv.Logger)
	//if err != nil {
	//	srv.Logger.Debug("register redis error", "error", err)
	//	return false
	//}
	err = rbac.InitGorm(srv.DB)
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.ModelProvider.Register("enforcer", rbac.GetEnforcer())

	return srv.registerDatabaseService()
}

func (srv *Server) MockDatabase() bool {
	srv.Cfg.DatabaseConfig.Escaper = "\""
	var err error
	srv.DB, err = model.MockORM(srv.Cfg)
	if err != nil {
		srv.Logger.Error("open database error", "error", err)
		return false
	}
	srv.Module.Provide(config.ModulePath.DBInstance.GormDB, srv.DB)

	if !model.Register(srv.Module) {
		return false
	}

	err = rbac.InitGorm(srv.DB)
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.ModelProvider.Register("enforcer", rbac.GetEnforcer())

	return srv.registerDatabaseService()
}
