package model

import (
	"github.com/Myriad-Dreamin/ginx/config"
	splayer "github.com/Myriad-Dreamin/ginx/model/sp-layer"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gomodule/redigo/redis"
	"time"
)

func OpenRedis(cfg *config.ServerConfig) (*redis.Pool, error) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				cfg.RedisConfig.ConnectionType, cfg.RedisConfig.Host,
				redis.DialPassword(cfg.RedisConfig.Password),
				redis.DialDatabase(cfg.RedisConfig.Database),
				redis.DialConnectTimeout(cfg.RedisConfig.ConnectionTimeout),
				redis.DialReadTimeout(cfg.RedisConfig.ReadTimeout),
				redis.DialWriteTimeout(cfg.RedisConfig.WriteTimeout),
				redis.DialKeepAlive(time.Minute*5),
			)
		},
		//TestOnBorrow:    nil,
		MaxIdle:     cfg.RedisConfig.MaxIdle,
		MaxActive:   cfg.RedisConfig.MaxActive,
		IdleTimeout: cfg.RedisConfig.IdleTimeout,
		Wait:        cfg.RedisConfig.Wait,
		//MaxConnLifetime: 0,
	}
	return pool, nil
}

func RegisterRedis(pool *redis.Pool, logger types.Logger) error {
	return splayer.RegisterRedis(pool, logger)
}

