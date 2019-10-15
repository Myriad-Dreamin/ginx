package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (srv *Service) Delete(c *gin.Context) {
	obj := new(model.Object)
	var ok bool
	obj.ID, ok = ginhelper.ParseUint(c, "oid")
	if !ok {
		return
	}
	if ginhelper.DeleteObj(c, obj) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}

