package base_service


import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"net/http"
)


type CObjectToolLite interface {
	FObject
	CObject
}

type CServiceInterface interface {
	Post(c controller.MContext)
}

type CService struct {
	Tool CObjectToolLite
	k    string
}

func NewCService(tool CObjectToolLite, k string) CService {
	return CService{
		Tool: tool,
		k:    k,
	}
}

func (srv CService) Post(c controller.MContext) {
	var obj = srv.Tool.SerializePost(c)
	if c.IsAborted() {
		return
	}
	if ginhelper.CreateObj(c, obj) {
		c.JSON(http.StatusOK, srv.Tool.ResponsePost(obj))
	}
}