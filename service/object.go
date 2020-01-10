package service

import (
	"github.com/Myriad-Dreamin/minimum-template/control"
	objectservice "github.com/Myriad-Dreamin/minimum-template/service/object"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type ObjectService = control.ObjectService

func NewObjectService(m module.Module) (ObjectService, error) {
	return objectservice.NewService(m)
}

func (s *Provider) ObjectService() ObjectService {
	return s.objectService
}
