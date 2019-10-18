package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
)

func (srv *Service) CreateObject(id uint) base_service.CRUDObject {
	return &model.Object{ID: id}
}

func (srv *Service) GetObject(id uint) (base_service.CRUDObject, error) {
	return srv.db.ID(id)
}

func (srv *Service) ResponsePost(obj base_service.CRUDObject) interface{} {
	return ObjectToPostReply(obj.(*model.Object))
}

func (srv *Service) ResponseGet(obj base_service.CRUDObject) interface{} {
	return ObjectToGetReply(obj.(*model.Object))
}

func (srv *Service) GetPutRequest() interface{} {
	return new(PutRequest)
}

func (srv *Service) FillPutFields(object base_service.CRUDObject, req interface{}) []string {
	return srv.fillPutFields(object.(*model.Object), req.(*PutRequest))
}
