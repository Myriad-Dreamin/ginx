package types

import (
	"github.com/gin-gonic/gin"
)

type Hook struct {
	funcs []GinHookFunc
}

type GinHookFunc func (c *gin.Context) bool

func (hook *Hook) Use(hookFunc GinHookFunc) {
	hook.funcs = append(hook.funcs, hookFunc)
}

func (hook *Hook) Consume(c *gin.Context) bool {
	for _, f := range hook.funcs {
		if !f(c) {
			return false
		}
	}
	return true
}
