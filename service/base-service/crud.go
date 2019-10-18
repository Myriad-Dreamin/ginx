package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)



type CRUDObject interface {
	Create() (int64, error)
	UpdateFields(fields []string) (int64, error)
	Delete() (int64, error)
}

type CRUDObjectToolLite interface {
	CreateObject(id uint) CRUDObject
	GetObject(id uint) (CRUDObject, error)
	SerializePost(c *gin.Context) CRUDObject
	ResponsePost(obj CRUDObject) interface{}
	ResponseGet(obj CRUDObject) interface{}
	GetPutRequest() interface{}
	FillPutFields(object CRUDObject, req interface{}) []string
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
	if ginhelper.DeleteObj(c, srv.Tool.CreateObject(id)) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}


func (srv *CRUDService) Get(c *gin.Context) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetObject(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, srv.Tool.ResponseGet(obj))
}


func (srv *CRUDService) Put(c *gin.Context) {
	var req =  srv.Tool.GetPutRequest()
	id, ok := ginhelper.ParseUintAndBind(c, srv.k, req)
	if !ok {
		return
	}

	object, err := srv.Tool.GetObject(id)
	if ginhelper.MaybeSelectError(c, object, err) {
		return
	}

	if ginhelper.UpdateFields(c, object, srv.Tool.FillPutFields(object, req)) {
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



