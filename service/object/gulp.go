package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control"
	base_service "github.com/Myriad-Dreamin/minimum-template/lib/base-service"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

func (svc *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Object{ID: id}
}

func (svc *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return svc.db.ID(id)
}

func (svc *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return svc.AfterPost(control.SerializePostObjectReply(types.CodeOK, obj.(*model.Object)))
}

func (svc *Service) DeleteHook(c controller.MContext, obj base_service.CRUDEntity) bool {
	return svc.deleteHook(c, obj.(*model.Object))
}

func (svc *Service) ResponseGet(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeGetObjectReply(types.CodeOK, obj.(*model.Object))
}

func (svc *Service) ResponseInspect(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeInspectObjectReply(types.CodeOK, obj.(*model.Object))
}

func (svc *Service) ProcessListResults(_ controller.MContext, result interface{}) interface{} {
	return control.PSerializeListObjectsReply(types.CodeOK, result.([]model.Object))
}

func (svc *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (svc *Service) GetPutRequest() interface{} {
	return new(control.PutObjectRequest)
}

func (svc *Service) FillPutFields(c controller.MContext, object base_service.CRUDEntity, req interface{}) []string {
	return svc.fillPutFields(c, object.(*model.Object), req.(*control.PutObjectRequest))
}

func (svc *Service) ListObjects(c controller.MContext) {
	svc.List(c)
	return
}

func (svc *Service) GetObject(c controller.MContext) {
	svc.Get(c)
	return
}

func (svc *Service) PutObject(c controller.MContext) {
	svc.Put(c)
	return
}

func (svc *Service) PostObject(c controller.MContext) {
	svc.Post(c)
	return
}

func (svc *Service) InspectObject(c controller.MContext) {
	svc.Inspect(c)
	return
}
