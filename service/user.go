package service

import (
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/ginx/config"
	"github.com/Myriad-Dreamin/ginx/control"
	"github.com/Myriad-Dreamin/ginx/model"
	userservice "github.com/Myriad-Dreamin/ginx/service/user"
	"github.com/Myriad-Dreamin/ginx/types"
)

type UserService = control.UserService

func NewUserService(logger types.Logger, provider *model.Provider, middleware *jwt.Middleware, config *config.ServerConfig) (UserService, error) {
	return userservice.NewService(logger, provider, middleware, config)
}

func (s *Provider) UserService() UserService {
	return s.userService
}
