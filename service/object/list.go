package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (srv *Service) List(c *gin.Context) {
	page, pageSize, ok := ginhelper.RosolvePageVariable(c)
	if !ok {
		return
	}

	objs, err := srv.db.QueryChain().Page(page, pageSize).Query()
	if ginhelper.MaybeSelectError(c, objs, err) {
		return
	}

	c.JSON(http.StatusOK, ObjectsToListReply(objs))
	return
}
