package ginhelper

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResetPassword(c *gin.Context, obj *model.User, password string) bool {
	_, err := obj.ResetPassword(password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return false
	}
	return true
}

func AuthenticatePassword(c *gin.Context, user *model.User, password string) bool {
	if ok, err := user.AuthenticatePassword(password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeAuthenticatePasswordError,
			Error: err.Error(),
		})
		return false
	} else if !ok {
		c.JSON(http.StatusOK, &Response{
			Code: types.CodeUserWrongPassword,
		})
		return false
	}
	return true
}
