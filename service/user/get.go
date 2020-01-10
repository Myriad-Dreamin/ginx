package userservice

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"time"
)

type GetReply struct {
	Code      types.CodeType `json:"code"`
	NickName  string         `json:"nick_name" form:"nick_name"`
	LastLogin time.Time      `json:"last_login" form:"last_login"`
}

func UserToGetReply(obj *model.User) GetReply {
	return GetReply{
		Code:      types.CodeOK,
		NickName:  obj.NickName,
		LastLogin: obj.LastLogin,
	}
}
