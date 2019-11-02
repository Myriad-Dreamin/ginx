package userservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/types"
)

type GetReply struct {
	Code int         `json:"code"`
	User *model.User `json:"user"`
}

func UserToGetReply(obj *model.User) *GetReply {
	return &GetReply{
		Code: types.CodeOK,
		User: obj,
	}
}
