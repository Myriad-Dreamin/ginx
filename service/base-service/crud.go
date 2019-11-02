package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CRUDEntity interface {
	Create() (int64, error)
	UpdateFields(fields []string) (int64, error)
	Delete() (int64, error)
}

type CRUDObjectToolLite interface {
	CreateEntity(id uint) CRUDEntity
	GetEntity(id uint) (CRUDEntity, error)
	SerializePost(c *gin.Context) CRUDEntity
	DeleteHook(c *gin.Context, obj CRUDEntity) bool
	ResponsePost(obj CRUDEntity) interface{}
	ResponseGet(obj CRUDEntity) interface{}
	GetPutRequest() interface{}
	FillPutFields(c *gin.Context, object CRUDEntity, req interface{}) []string
}

type CRUDService struct {
	Tool CRUDObjectToolLite
	k    string
}

func NewCRUDService(tool CRUDObjectToolLite, k string) CRUDService {
	return CRUDService{
		Tool: tool,
		k:    k,
	}
}

func (srv *CRUDService) Delete(c *gin.Context) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}
	if !srv.Tool.DeleteHook(c, obj) {
		return
	}

	if ginhelper.DeleteObj(c, obj) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}

func (srv *CRUDService) Get(c *gin.Context) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, srv.Tool.ResponseGet(obj))
}

func (srv *CRUDService) Put(c *gin.Context) {
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

func (srv *CRUDService) Post(c *gin.Context) {
	var obj = srv.Tool.SerializePost(c)
	if c.IsAborted() {
		return
	}
	if ginhelper.CreateObj(c, obj) {
		c.JSON(http.StatusOK, srv.Tool.ResponsePost(obj))
	}
}
