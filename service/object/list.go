package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
)

type ListReply struct {
	Code    int            `json:"code"`
	Objects []model.Object `json:"objects"`
}

func ObjectsToListReply(obj []model.Object) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Objects = obj
	return
}

func (srv *Service) FilterOn(c *gin.Context) (interface{}, error) {
	// parse c

	page, pageSize, ok := ginhelper.RosolvePageVariable(c)
	if !ok {
		return nil, nil
	}

	objs, err := srv.db.QueryChain().Page(page, pageSize).Query()
	if err != nil {
		return nil, err
	}
	return ObjectsToListReply(objs), nil
}
