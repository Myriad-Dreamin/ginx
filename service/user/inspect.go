package userservice

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

type InspectReply struct {
	Code types.CodeType `json:"code"`
	User *model.User    `json:"user"`
}

func UserToInspectReply(obj *model.User) *InspectReply {
	return &InspectReply{
		Code: types.CodeOK,
		User: obj,
	}
}
