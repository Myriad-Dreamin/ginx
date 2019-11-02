package userservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
	"github.com/gin-gonic/gin"
)

func (srv *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.User{ID: id}
}

func (srv *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return srv.db.ID(id)
}

func (srv *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return UserToPostReply(obj.(*model.User))
}

func (srv *Service) DeleteHook(c *gin.Context, obj base_service.CRUDEntity) bool {
	return srv.deleteHook(c, obj.(*model.User))
}

func (srv *Service) ResponseGet(obj base_service.CRUDEntity) interface{} {
	return UserToGetReply(obj.(*model.User))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(c *gin.Context, user base_service.CRUDEntity, req interface{}) []string {
	return srv.fillPutFields(c, user.(*model.User), req.(*PutRequest))
}
