package userservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/gin-gonic/gin"
)

func (srv *Service) deleteHook(c *gin.Context, object *model.User) (canDelete bool) {
	return false
}
