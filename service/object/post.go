package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostRequest struct {
}

type PostReply struct {
	Code   int           `form:"code" json:"code"`
	Object *model.Object `form:"object" json:"object"`
}

func (srv *Service) Post(c *gin.Context) {
	var req = new(PostRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	var obj = new(model.Object)
	if ginhelper.CreateObj(c, obj) {
		//cr.Submissionr.PushTask(code)
		c.JSON(http.StatusOK, &PostReply{
			Code:   types.CodeOK,
			Object: obj,
		})
	}
}
