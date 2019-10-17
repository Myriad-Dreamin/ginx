package service

import (
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/model"
	objectservice "github.com/Myriad-Dreamin/ginx/service/object"
	"github.com/Myriad-Dreamin/ginx/types"
)

type ObjectService = objectservice.Service

func NewObjectService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (*ObjectService, error) {
	return objectservice.NewService(logger, provider, config)
}

func (s *Provider) ObjectService() *ObjectService {
	return s.objectService
}
