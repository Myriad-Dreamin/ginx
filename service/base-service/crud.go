package base_service

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"net/http"
)

type CRUDObjectToolLite interface {
	FObject
	DObject
	SerializePost(c controller.MContext) CRUDEntity
	ResponsePost(obj CRUDEntity) interface{}
	ResponseGet(c controller.MContext, obj CRUDEntity) interface{}
	ResponseInspect(obj CRUDEntity) interface{}
	GetPutRequest() interface{}
	FillPutFields(c controller.MContext, object CRUDEntity, req interface{}) []string
}

type CRUDService struct {
	Tool CRUDObjectToolLite
	k    string
	DServiceInterface
}

func NewCRUDService(tool CRUDObjectToolLite, k string) CRUDService {
	return CRUDService{
		Tool:              tool,
		k:                 k,
		DServiceInterface: NewDService(tool, k),
	}
}

func (srv *CRUDService) Get(c controller.MContext) {
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

func (srv *CRUDService) Inspect(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, srv.Tool.ResponseInspect(obj))
}

func (srv *CRUDService) Put(c controller.MContext) {
	var req = srv.Tool.GetPutRequest()
	id, ok := ginhelper.ParseUintAndBind(c, srv.k, req)
	if !ok {
		return
	}

	object, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, object, err) {
		return
	}

	fields := srv.Tool.FillPutFields(c, object, req)
	if c.IsAborted() {
		return
	}

	if ginhelper.UpdateFields(c, object, fields) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}

func (srv *CRUDService) Post(c controller.MContext) {
	var obj = srv.Tool.SerializePost(c)
	if c.IsAborted() {
		return
	}
	if ginhelper.CreateObj(c, obj) {
		c.JSON(http.StatusOK, srv.Tool.ResponsePost(obj))
	}
}
