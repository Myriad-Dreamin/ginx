package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control"
	base_service "github.com/Myriad-Dreamin/minimum-template/lib/base-service"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
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
	panic("abort")
	//return UserToPostReply(obj.(*model.User))
}

func (srv *Service) DeleteHook(c controller.MContext, obj base_service.CRUDEntity) bool {
	return srv.deleteHook(c, obj.(*model.User))
}

func (srv *Service) ResponseGet(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeGetUserReply(types.CodeOK, obj.(*model.User))
}

func (srv *Service) ResponseInspect(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeInspectUserReply(types.CodeOK, obj.(*model.User))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(control.PutUserRequest)
}

func (srv *Service) FillPutFields(c controller.MContext, user base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(c, user.(*model.User), req.(*control.PutUserRequest))
}


func (srv *Service) ListUsers(c controller.MContext) {
	srv.List(c)
	return
}

func (srv *Service) InspectUser(c controller.MContext) {
	srv.Inspect(c)
	return
}

func (srv *Service) GetUser(c controller.MContext) {
	srv.Get(c)
	return
}

func (srv *Service) PutUser(c controller.MContext) {
	srv.Put(c)
	return
}

