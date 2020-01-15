package control

import (
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/model/db-layer"
)

type ObjectService interface {
	ObjectServiceSignatureXXX() interface{}
	ListObjects(c controller.MContext)
	PostObject(c controller.MContext)
	InspectObject(c controller.MContext)
	GetObject(c controller.MContext)
	PutObject(c controller.MContext)
	Delete(c controller.MContext)
}
type ListObjectsRequest = gorm_crud_dao.Filter

type ListObjectsReply struct {
	Code    int              `json:"code" form:"code"`
	Objects []dblayer.Object `json:"objects" form:"objects"`
}

type PostObjectRequest struct {
}

type PostObjectReply struct {
	Code   int             `json:"code" form:"code"`
	Object *dblayer.Object `json:"object" form:"object"`
}

type InspectObjectReply struct {
	Code   int             `json:"code" form:"code"`
	Object *dblayer.Object `json:"object" form:"object"`
}

type GetObjectReply struct {
	Code   int             `json:"code" form:"code"`
	Object *dblayer.Object `json:"object" form:"object"`
}

type PutObjectRequest struct {
}

func PSerializeListObjectsReply(_code int, _objects []dblayer.Object) *ListObjectsReply {

	return &ListObjectsReply{
		Code:    _code,
		Objects: _objects,
	}
}
func SerializeListObjectsReply(_code int, _objects []dblayer.Object) ListObjectsReply {

	return ListObjectsReply{
		Code:    _code,
		Objects: _objects,
	}
}
func _packSerializeListObjectsReply(_code int, _objects []dblayer.Object) ListObjectsReply {

	return ListObjectsReply{
		Code:    _code,
		Objects: _objects,
	}
}
func PackSerializeListObjectsReply(_code []int, _objects [][]dblayer.Object) (pack []ListObjectsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListObjectsReply(_code[i], _objects[i]))
	}
	return
}
func PSerializePostObjectRequest() *PostObjectRequest {

	return &PostObjectRequest{}
}
func SerializePostObjectRequest() PostObjectRequest {

	return PostObjectRequest{}
}
func _packSerializePostObjectRequest() PostObjectRequest {

	return PostObjectRequest{}
}
func PackSerializePostObjectRequest() (pack []PostObjectRequest) {
	return
}
func PSerializePostObjectReply(_code int, _object *dblayer.Object) *PostObjectReply {

	return &PostObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func SerializePostObjectReply(_code int, _object *dblayer.Object) PostObjectReply {

	return PostObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func _packSerializePostObjectReply(_code int, _object *dblayer.Object) PostObjectReply {

	return PostObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func PackSerializePostObjectReply(_code []int, _object []*dblayer.Object) (pack []PostObjectReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostObjectReply(_code[i], _object[i]))
	}
	return
}
func PSerializeInspectObjectReply(_code int, _object *dblayer.Object) *InspectObjectReply {

	return &InspectObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func SerializeInspectObjectReply(_code int, _object *dblayer.Object) InspectObjectReply {

	return InspectObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func _packSerializeInspectObjectReply(_code int, _object *dblayer.Object) InspectObjectReply {

	return InspectObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func PackSerializeInspectObjectReply(_code []int, _object []*dblayer.Object) (pack []InspectObjectReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectObjectReply(_code[i], _object[i]))
	}
	return
}
func PSerializeGetObjectReply(_code int, _object *dblayer.Object) *GetObjectReply {

	return &GetObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func SerializeGetObjectReply(_code int, _object *dblayer.Object) GetObjectReply {

	return GetObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func _packSerializeGetObjectReply(_code int, _object *dblayer.Object) GetObjectReply {

	return GetObjectReply{
		Code:   _code,
		Object: _object,
	}
}
func PackSerializeGetObjectReply(_code []int, _object []*dblayer.Object) (pack []GetObjectReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetObjectReply(_code[i], _object[i]))
	}
	return
}
func PSerializePutObjectRequest() *PutObjectRequest {

	return &PutObjectRequest{}
}
func SerializePutObjectRequest() PutObjectRequest {

	return PutObjectRequest{}
}
func _packSerializePutObjectRequest() PutObjectRequest {

	return PutObjectRequest{}
}
func PackSerializePutObjectRequest() (pack []PutObjectRequest) {
	return
}
