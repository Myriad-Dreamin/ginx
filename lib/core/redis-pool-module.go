package mcore

import (
	"time"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gomodule/redigo/redis"
)

type RedisPoolModule struct {
	RedisPool *redis.Pool
}

func (m *RedisPoolModule) FromRaw(db *redis.Pool, dep module.Module) bool {
	m.RedisPool = db
	dep.Provide(DefaultNamespace.DBInstance.RedisPool, db)
	return true
}

func (m *RedisPoolModule) FromContext(dep module.Module) bool {
	m.RedisPool = dep.Require(DefaultNamespace.DBInstance.RedisPool).(*redis.Pool)
	return true
}

func (m *RedisPoolModule) Install(dep module.Module) bool {
	return m.FromContext(dep)
}

func (m *RedisPoolModule) InstallFromConfiguration(dep module.Module) bool {
	xdb, err := OpenRedis(dep)
	m.FromRaw(xdb, dep)
	return Maybe(dep, "init redis error", err)}

func (m *RedisPoolModule) GetRedisPoolInstance() *redis.Pool {
	return m.RedisPool
}

func OpenRedis(dep module.Module) (*redis.Pool, error) {
	cfg := getRedisConfiguration(dep)
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				cfg.ConnectionType, cfg.Host,
				redis.DialPassword(cfg.Password),
				redis.DialDatabase(cfg.Database),
				redis.DialConnectTimeout(cfg.ConnectionTimeout),
				redis.DialReadTimeout(cfg.ReadTimeout),
				redis.DialWriteTimeout(cfg.WriteTimeout),
				redis.DialKeepAlive(time.Minute*5),
			)
		},
		//TestOnBorrow:    nil,
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout,
		Wait:        cfg.Wait,
		//MaxConnLifetime: 0,
	}
	return pool, nil
}

