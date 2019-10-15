package service

import (
	"fmt"
	"github.com/Myriad-Dreamin/ginx/service/object"
)

type Provider struct {
	objectService *objectservice.Service
}

func (s *Provider) Register(service interface{}) {
	switch ss := service.(type) {
	case *objectservice.Service:
		s.objectService = ss
	default:
		panic(fmt.Errorf("unknown service %T", service))
	}
}

func (s *Provider) ObjectService() *objectservice.Service {
	return s.objectService
}
