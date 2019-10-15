package model

import (
	splayer "github.com/Myriad-Dreamin/ginx/model/sp-layer"
	"github.com/Myriad-Dreamin/ginx/types"
)

type Object = splayer.Object
type ObjectDB = splayer.ObjectDB

func NewObjectDB(logger types.Logger) (*ObjectDB, error) {
	return splayer.NewObjectDB(logger)
}

func GetObjectDB(logger types.Logger) (*ObjectDB, error) {
	return splayer.GetObjectDB(logger)
}