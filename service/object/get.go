package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

type GetReply struct {
	Code   int           `json:"code"`
	Object *model.Object `json:"object"`
}

func ObjectToGetReply(obj *model.Object) *GetReply {
	return &GetReply{
		Code:   types.CodeOK,
		Object: obj,
	}
}
