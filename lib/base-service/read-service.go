package base_service

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"net/http"
)

type RObjectToolLite interface {
	FObject
	RObject
}

type RServiceInterface interface {
	Get(c controller.MContext)
}

type RService struct {
	Tool RObjectToolLite
	k    string
}

func NewRService(tool RObjectToolLite, k string) RService {
	return RService{
		Tool: tool,
		k:    k,
	}
}

func (srv RService) Get(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}
	x := srv.Tool.ResponseGet(c, obj)
	if c.IsAborted() {
		return
	}
	c.JSON(http.StatusOK, x)
}

func (srv RService) Inspect(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	x := srv.Tool.ResponseInspect(c, obj)
	if c.IsAborted() {
		return
	}
	c.JSON(http.StatusOK, x)
}