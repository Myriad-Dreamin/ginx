package mgin

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (c Context) GetMeta() interface{} {
	v, _ := c.Get("__meta")
	return v
}

func (c Context) SetMeta(meta interface{}) {
	c.Set("__meta", meta)
}

func (c Context) Copy() controller.MContext {
	return &Context{c.Context.Copy()}
}

func (c Context) AbortWithError(code int, err error) controller.Error {
	return GinError{c.Context.AbortWithError(code, err)}
}

func (c Context) Error(err error) controller.Error {
	return GinError{c.Context.Error(err)}
}
