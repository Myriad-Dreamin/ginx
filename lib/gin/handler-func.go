package mgin

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin"
)

type HandlerFunc = controller.HandlerFunc

func ToGinHandler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(Context{Context: c})
		return
	}
}

func ToGinHandlers(funcs []controller.HandlerFunc) (rFuncs []gin.HandlerFunc) {
	rFuncs = make([]gin.HandlerFunc, len(funcs))
	for i := range funcs {
		rFuncs[i] = ToGinHandler(funcs[i])
	}
	return
}
