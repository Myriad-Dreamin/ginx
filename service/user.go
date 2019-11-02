package service

import (
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/control"
	"github.com/Myriad-Dreamin/minimum-template/model"
	userservice "github.com/Myriad-Dreamin/minimum-template/service/user"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

type UserService = control.UserService

func NewUserService(logger types.Logger, provider *model.Provider, middleware *jwt.Middleware, config *config.ServerConfig) (UserService, error) {
	return userservice.NewService(logger, provider, middleware, config)
}

func (s *Provider) UserService() UserService {
	return s.userService
}
