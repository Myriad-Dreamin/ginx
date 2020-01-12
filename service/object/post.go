package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	base_service "github.com/Myriad-Dreamin/minimum-template/lib/base-service"
	"github.com/Myriad-Dreamin/minimum-template/model"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

type PostReply struct {
	Code   types.CodeRawType `json:"code"`
	Object *model.Object  `json:"object"`
}

func ObjectToPostReply(obj *model.Object) *PostReply {
	return &PostReply{
		Code:   types.CodeOK,
		Object: obj,
	}
}

type PostRequest struct {
}

func (srv *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Object)
	// fill here
	return obj
}

func (srv *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
