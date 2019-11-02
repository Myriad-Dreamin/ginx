package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/types"
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
