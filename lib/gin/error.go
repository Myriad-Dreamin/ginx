package mgin

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
)

type GinError struct {
	*gin.Error
}

func (g GinError) GetError() error {
	return g.Err
}

func (g GinError) GetType() controller.ErrorType {
	return uint64(g.Type)
}

func (g GinError) GetMeta() interface{} {
	return g.Meta
}
