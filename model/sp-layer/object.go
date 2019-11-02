package splayer

import (
	"github.com/Myriad-Dreamin/minimum-template/config"
	dblayer "github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

type Object = dblayer.Object

//struct {
//	dblayer.Object
//}

//func (c *Object) Update() (int64, error) {
//	return c.Object.Update()
//}

//type _userConverter struct {}
//var ObjectConverter _userConverter
//func (_userConverter) convert(obj dblayer.Object) Object { return Object{Object: obj} }
//func (_userConverter) reconvert(obj Object) dblayer.Object { return obj.Object }
//func (_userConverter) converts(objs []dblayer.Object) []Object { return *(*[]Object)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconverts(objs []Object) []dblayer.Object { return *(*[]dblayer.Object)(unsafe.Pointer(&objs)) }
//func (_userConverter) convertP(obj *dblayer.Object) *Object { return (*Object)(unsafe.Pointer(obj)) }
//func (_userConverter) reconvertP(obj *Object) *dblayer.Object { return (*dblayer.Object)(unsafe.Pointer(obj)) }
//func (_userConverter) convertPs(objs []*dblayer.Object) []*Object { return *(*[]*Object)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconvertPs(objs []*Object) []*dblayer.Object { return *(*[]*dblayer.Object)(unsafe.Pointer(&objs)) }

type ObjectDB struct {
	dblayer.ObjectDB
}

func NewObjectDB(logger types.Logger, cfg *config.ServerConfig) (*ObjectDB, error) {
	cdb, err := dblayer.NewObjectDB(logger, cfg)
	if err != nil {
		return nil, err
	}
	db := new(ObjectDB)
	db.ObjectDB = *cdb
	return db, nil
}

func GetObjectDB(logger types.Logger, cfg *config.ServerConfig) (*ObjectDB, error) {
	cdb, err := dblayer.GetObjectDB(logger, cfg)
	if err != nil {
		return nil, err
	}
	db := new(ObjectDB)
	db.ObjectDB = *cdb
	return db, nil
}

func (s *Provider) ObjectDB() *ObjectDB {
	return s.objectDB
}
