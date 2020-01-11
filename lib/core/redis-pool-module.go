package mcore

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gomodule/redigo/redis"
)

type RedisPoolModule struct {
	redisPool *redis.Pool
}

func (m *RedisPoolModule) FromRaw(db *redis.Pool, dep module.Module) bool {
	m.redisPool = db
	dep.Provide(DefaultNamespace.DBInstance.RedisPool, db)
	return true
}

func (m *RedisPoolModule) FromContext(dep module.Module) bool {
	m.redisPool = dep.Require(DefaultNamespace.DBInstance.RedisPool).(*redis.Pool)
	return true
}

func (m *RedisPoolModule) Install(dep module.Module) bool {
	return m.FromContext(dep)
}

func (m *RedisPoolModule) GetRedisPoolInstance() *redis.Pool {
	return m.redisPool
}
