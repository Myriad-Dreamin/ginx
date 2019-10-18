package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListableObjectToolLite interface {
	FilterOn(c *gin.Context, page, pageSize int) (interface{}, error)
}

type ListService struct {
	tool ListableObjectToolLite
	k string
}

func NewListService(tool ListableObjectToolLite, k string) ListService {
	return ListService{
		tool: tool,
		k:   k,
	}
}


func (srv *ListService) List(c *gin.Context) {
	page, pageSize, ok := ginhelper.RosolvePageVariable(c)
	if !ok {
		return
	}

	result, err := srv.tool.FilterOn(c, page, pageSize)
	if c.IsAborted() || ginhelper.MaybeOnlySelectError(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

