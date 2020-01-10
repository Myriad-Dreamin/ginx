package splayer

import (
	dblayer "github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type User = dblayer.User


type UserDB struct {
	dblayer.UserDB
}

func NewUserDB(m module.Module) (*UserDB, error) {
	cdb, err := dblayer.NewUserDB(m)
	if err != nil {
		return nil, err
	}
	db := new(UserDB)
	db.UserDB = *cdb
	return db, nil
}

func GetUserDB(m module.Module) (*UserDB, error) {
	cdb, err := dblayer.GetUserDB(m)
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
