package model

import (
	splayer "github.com/Myriad-Dreamin/minimum-template/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Object = splayer.Object
type ObjectDB = splayer.ObjectDB

func NewObjectDB(m module.Module) (*ObjectDB, error) {
	return splayer.NewObjectDB(m)
}

func GetObjectDB(m module.Module) (*ObjectDB, error) {
	return splayer.GetObjectDB(m)
}
