package ginhelper

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"net/http"
)

func ResetPassword(c controller.MContext, obj *model.User, password string) bool {
	_, err := obj.ResetPassword(password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return false
	}
	return true
}

func AuthenticatePassword(c controller.MContext, user *model.User, password string) bool {
	if ok, err := user.AuthenticatePassword(password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeAuthenticatePasswordError,
			Error: err.Error(),
		})
		return false
	} else if !ok {
		c.JSON(http.StatusOK, &types.Response{
			Code: types.CodeUserWrongPassword,
		})
		return false
	}
	return true
}
