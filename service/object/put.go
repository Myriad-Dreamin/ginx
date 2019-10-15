package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PutRequest struct {
}

func (srv *Service) Put(c *gin.Context) {
	var req =  new(PutRequest)
	id, ok := ginhelper.ParseUintAndBind(c, "oid", req)
	if !ok {
		return
	}

	announcement, err := srv.db.ID(id)
	if ginhelper.MaybeSelectError(c, announcement, err) {
		return
	}

	if ginhelper.UpdateFields(c, announcement, srv.FillPutFields(announcement, req)) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}


func (srv *Service) FillPutFields(announcement *model.Object, req *PutRequest) (fields []string) {
	return
}