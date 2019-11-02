package userservice

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
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
