package splayer

import (
	"github.com/Myriad-Dreamin/ginx/config"
	dblayer "github.com/Myriad-Dreamin/ginx/model/db-layer"
	"github.com/Myriad-Dreamin/ginx/types"
)

type User = dblayer.User

//struct {
//	dblayer.User
//}

//func (c *User) Update() (int64, error) {
//	return c.User.Update()
//}

//type _userConverter struct {}
//var UserConverter _userConverter
//func (_userConverter) convert(obj dblayer.User) User { return User{User: obj} }
//func (_userConverter) reconvert(obj User) dblayer.User { return obj.User }
//func (_userConverter) converts(objs []dblayer.User) []User { return *(*[]User)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconverts(objs []User) []dblayer.User { return *(*[]dblayer.User)(unsafe.Pointer(&objs)) }
//func (_userConverter) convertP(obj *dblayer.User) *User { return (*User)(unsafe.Pointer(obj)) }
//func (_userConverter) reconvertP(obj *User) *dblayer.User { return (*dblayer.User)(unsafe.Pointer(obj)) }
//func (_userConverter) convertPs(objs []*dblayer.User) []*User { return *(*[]*User)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconvertPs(objs []*User) []*dblayer.User { return *(*[]*dblayer.User)(unsafe.Pointer(&objs)) }

type UserDB struct {
	dblayer.UserDB
}

func NewUserDB(logger types.Logger, cfg *config.ServerConfig) (*UserDB, error) {
	cdb, err := dblayer.NewUserDB(logger, cfg)
	if err != nil {
		return nil, err
	}
	db := new(UserDB)
	db.UserDB = *cdb
	return db, nil
}

func GetUserDB(logger types.Logger, cfg *config.ServerConfig) (*UserDB, error) {
	cdb, err := dblayer.GetUserDB(logger, cfg)
	if err != nil {
		return nil, err
	}
	db := new(UserDB)
	db.UserDB = *cdb
	return db, nil
}

func (s *Provider) UserDB() *UserDB {
	return s.userDB
}
