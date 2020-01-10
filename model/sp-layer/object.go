package splayer

import (
	dblayer "github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Object = dblayer.Object


type ObjectDB struct {
	dblayer.ObjectDB
}

func NewObjectDB(m module.Module) (*ObjectDB, error) {
	cdb, err := dblayer.NewObjectDB(m)
	if err != nil {
		return nil, err
	}
	db := new(ObjectDB)
	db.ObjectDB = *cdb
	return db, nil
}

func GetObjectDB(m module.Module) (*ObjectDB, error) {
	cdb, err := dblayer.GetObjectDB(m)
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
