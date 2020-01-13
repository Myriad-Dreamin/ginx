
package control

import (
    "github.com/Myriad-Dreamin/minimum-template/model/db-layer"
    "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
    "github.com/Myriad-Dreamin/minimum-lib/controller"

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
    Code int`json:"code"form:"code"`
    Objects []dblayer.Object`json:"objects"form:"objects"`
}

func PSerializeListObjectsReply(_code int, _objects []dblayer.Object) *ListObjectsReply {
	return &ListObjectsReply{
        Code: _code,
        Objects: _objects,
	}
}

func SerializeListObjectsReply(_code int, _objects []dblayer.Object) ListObjectsReply {
	return ListObjectsReply{
        Code: _code,
        Objects: _objects,
	}
}

func PackSerializeListObjectsReply(_code []int, _objects [][]dblayer.Object) (pack []ListObjectsReply) {
	for i := range _code {
		pack = append(pack, SerializeListObjectsReply(_code[i], _objects[i]))
	}
	return
}


type PostObjectRequest struct {

}

func PSerializePostObjectRequest() *PostObjectRequest {
	return &PostObjectRequest{

	}
}

func SerializePostObjectRequest() PostObjectRequest {
	return PostObjectRequest{

	}
}

func PackSerializePostObjectRequest() (pack []PostObjectRequest) {
	return
}


type PostObjectReply struct {
    Code int`json:"code"form:"code"`
    Object *dblayer.Object`json:"object"form:"object"`
}

func PSerializePostObjectReply(_code int, _object *dblayer.Object) *PostObjectReply {
	return &PostObjectReply{
        Code: _code,
        Object: _object,
	}
}

func SerializePostObjectReply(_code int, _object *dblayer.Object) PostObjectReply {
	return PostObjectReply{
        Code: _code,
        Object: _object,
	}
}

func PackSerializePostObjectReply(_code []int, _object []*dblayer.Object) (pack []PostObjectReply) {
	for i := range _code {
		pack = append(pack, SerializePostObjectReply(_code[i], _object[i]))
	}
	return
}


type InspectObjectReply struct {
    Code int`json:"code"form:"code"`
    User *dblayer.Object`json:"user"form:"user"`
}

func PSerializeInspectObjectReply(_code int, _user *dblayer.Object) *InspectObjectReply {
	return &InspectObjectReply{
        Code: _code,
        User: _user,
	}
}

func SerializeInspectObjectReply(_code int, _user *dblayer.Object) InspectObjectReply {
	return InspectObjectReply{
        Code: _code,
        User: _user,
	}
}

func PackSerializeInspectObjectReply(_code []int, _user []*dblayer.Object) (pack []InspectObjectReply) {
	for i := range _code {
		pack = append(pack, SerializeInspectObjectReply(_code[i], _user[i]))
	}
	return
}


type GetObjectReply struct {
    Code int`json:"code"form:"code"`
    Object *dblayer.Object`json:"object"form:"object"`
}

func PSerializeGetObjectReply(_code int, _object *dblayer.Object) *GetObjectReply {
	return &GetObjectReply{
        Code: _code,
        Object: _object,
	}
}

func SerializeGetObjectReply(_code int, _object *dblayer.Object) GetObjectReply {
	return GetObjectReply{
        Code: _code,
        Object: _object,
	}
}

func PackSerializeGetObjectReply(_code []int, _object []*dblayer.Object) (pack []GetObjectReply) {
	for i := range _code {
		pack = append(pack, SerializeGetObjectReply(_code[i], _object[i]))
	}
	return
}


type PutObjectRequest struct {

}

func PSerializePutObjectRequest() *PutObjectRequest {
	return &PutObjectRequest{

	}
}

func SerializePutObjectRequest() PutObjectRequest {
	return PutObjectRequest{

	}
}

func PackSerializePutObjectRequest() (pack []PutObjectRequest) {
	return
}
