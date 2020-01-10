package base_service

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

type CRUDEntity interface {
	Create() (int64, error)
	UpdateFields(fields []string) (int64, error)
	Delete() (int64, error)
}

type dHookObject interface {
	DeleteHook(c controller.MContext, obj CRUDEntity) bool
}

type FObject interface {
	CreateEntity(id uint) CRUDEntity
	GetEntity(id uint) (CRUDEntity, error)
}

type DObject interface {
	dHookObject
}

type RObject interface {
	ResponseGet(c controller.MContext, obj CRUDEntity) interface{}
	ResponseInspect(c controller.MContext, obj CRUDEntity) interface{}
}

type UObject interface {
	GetPutRequest() interface{}
	FillPutFields(c controller.MContext, object CRUDEntity, req interface{}) []string
}
type CObject interface {
	SerializePost(c controller.MContext) CRUDEntity
	ResponsePost(obj CRUDEntity) interface{}
}
