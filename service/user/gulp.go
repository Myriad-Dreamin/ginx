package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	base_service "github.com/Myriad-Dreamin/minimum-template/lib/base-service"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.User{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.userDB.ID(id)
}

func (srv *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return UserToPostReply(obj.(*model.User))
}

func (srv *Service) DeleteHook(c controller.MContext, obj base_service.CRUDEntity) bool {
	return srv.deleteHook(c, obj.(*model.User))
}

func (srv *Service) ResponseGet(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return UserToGetReply(obj.(*model.User))
}

func (srv *Service) ResponseInspect(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return UserToInspectReply(obj.(*model.User))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(c controller.MContext, user base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(c, user.(*model.User), req.(*PutRequest))
}
