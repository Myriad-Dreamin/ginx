package service

import (
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/model"
	objectservice "github.com/Myriad-Dreamin/minimum-template/service/object"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type ObjectService = objectservice.Service

func NewObjectService(logger types.Logger, provider *model.Provider, config *config.ServerConfig) (*ObjectService, error) {
	return objectservice.NewService(logger, provider, config)
}

func (s *Provider) ObjectService() *ObjectService {
	return s.objectService
}
