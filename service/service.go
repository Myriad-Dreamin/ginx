package service

import (
	"github.com/Myriad-Dreamin/ginx/model"
	objectservice "github.com/Myriad-Dreamin/ginx/service/object"
	"github.com/Myriad-Dreamin/ginx/types"
)

func NewObjectService(logger types.Logger, provider *model.Provider) (*objectservice.Service, error) {
	return objectservice.NewService(logger, provider)
}
