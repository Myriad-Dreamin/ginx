package service

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-template/control"
	authservice "github.com/Myriad-Dreamin/minimum-template/service/auth"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type AuthService = control.AuthService

func NewAuthService(m module.Module) (AuthService, error) {
	return authservice.NewService(m)
}

func (s *Provider) AuthService() AuthService {
	return s.authService
}
