package userservice

import (
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChangePasswordRequest struct {
	OldPassword string `form:"old-password" json:"old-password" binding:"required,email"`
	NewPassword string `form:"new-password" json:"new-password" binding:"required,email"`
}

func (srv *Service) ChangePassword(c *gin.Context) {
	var req = new(ChangePasswordRequest)
	id, ok := ginhelper.ParseUintAndBind(c, "id", req)
	if !ok {
		return
	}

	user, err := srv.db.Query(id)
	if ginhelper.MaybeSelectError(c, user, err) {
		return
	}

	if !ginhelper.AuthenticatePassword(c, user, req.OldPassword) {
		return
	}

	if ginhelper.ResetPassword(c, user, req.NewPassword) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}
