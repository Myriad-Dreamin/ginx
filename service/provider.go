package service

import (
	"fmt"
)

type Provider struct {
	objectService *ObjectService
}

func (s *Provider) Register(service interface{}) {
	switch ss := service.(type) {
	case *ObjectService:
		s.objectService = ss
	default:
		panic(fmt.Errorf("unknown service %T", service))
	}
}
