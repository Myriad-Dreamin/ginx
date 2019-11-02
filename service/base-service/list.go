package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListableObjectToolLite interface {
	FilterOn(c *gin.Context) (interface{}, error)
}

type ListService struct {
	tool ListableObjectToolLite
	k    string
}

func NewListService(tool ListableObjectToolLite, k string) ListService {
	return ListService{
		tool: tool,
		k:    k,
	}
}

func (srv *ListService) List(c *gin.Context) {
	result, err := srv.tool.FilterOn(c)
	if c.IsAborted() || ginhelper.MaybeOnlySelectError(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
	return
}
