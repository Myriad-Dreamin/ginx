package model

import (
	"github.com/Myriad-Dreamin/ginx/config"
	splayer "github.com/Myriad-Dreamin/ginx/model/sp-layer"
	"github.com/Myriad-Dreamin/ginx/types"
)

type User = splayer.User
type UserDB = splayer.UserDB

func NewUserDB(logger types.Logger, cfg *config.ServerConfig) (*UserDB, error) {
	return splayer.NewUserDB(logger, cfg)
}

func GetUserDB(logger types.Logger, cfg *config.ServerConfig) (*UserDB, error) {
	return splayer.GetUserDB(logger, cfg)
}
