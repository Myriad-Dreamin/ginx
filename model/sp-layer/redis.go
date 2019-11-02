package splayer

import (
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool
var logger types.Logger

func RegisterRedis(mpool *redis.Pool, mlogger types.Logger) error {
	pool = mpool
	logger = mlogger
	return nil
}
