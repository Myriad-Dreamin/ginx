package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetReply struct {
	Code int `json:"code"`
	Object *model.Object
}

func ObjectToGetReply(obj *model.Object) *GetReply {
	return &GetReply{
		Code: types.CodeOK,
		Object: obj,
	}
}
func (srv *Service) Get(c *gin.Context) {
	id, ok := ginhelper.ParseUint(c, "oid")
	if !ok {
		return
	}
	obj, err := srv.db.ID(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	c.JSON(http.StatusOK, ObjectToGetReply(obj))
}
