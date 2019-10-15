package server

import (
	"github.com/Myriad-Dreamin/ginx/model"
)

func (srv *Server) registerDatabaseService() bool {

	objectDB, err := model.NewObjectDB(srv.Logger)
	if err != nil {
		srv.Logger.Debug("init object DB error", "error", err)
		return false
	}

	srv.DatabaseProvider.Register("objectdb", objectDB)
	return true
}

func (srv *Server) PrepareDatabase() bool {
	var err error
	cfg := srv.cfg
	srv.DB, err = model.OpenORM(cfg)
	if err != nil {
		srv.Logger.Debug("open database error", "error", err)
		return false
	}

	srv.Logger.Info("connected to database",
		"connection-type", cfg.DatabaseConfig.ConnectionType,
		"user", cfg.DatabaseConfig.User,
		"database", cfg.DatabaseConfig.DatabaseName,
		"charset", cfg.DatabaseConfig.Charset,
		"location", cfg.DatabaseConfig.Location,
	)

	err = model.Register(srv.DB, srv.Logger)
	if err != nil {
		srv.Logger.Error("register and migrate error", "error", err)
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

	return srv.registerDatabaseService()
}

