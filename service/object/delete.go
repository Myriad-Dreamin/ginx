package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/gin-gonic/gin"
)

func (srv *Service) deleteHook(c *gin.Context, object *model.Object) (canDelete bool) {
	return false
}
