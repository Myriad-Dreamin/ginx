package base_service


import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"net/http"
)


type UObjectToolLite interface {
	FObject
	UObject
}

type UServiceInterface interface {
	Put(c controller.MContext)
}

type UService struct {
	Tool UObjectToolLite
	k    string
}

func NewUService(tool UObjectToolLite, k string) UService {
	return UService{
		Tool: tool,
		k:    k,
	}
}

func (srv UService) Put(c controller.MContext) {
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
