package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control"
	base_service "github.com/Myriad-Dreamin/minimum-template/lib/base-service"
	"github.com/Myriad-Dreamin/minimum-template/model"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
)

func (svc *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req control.PostObjectRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Object)
	// fill here
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
