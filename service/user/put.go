package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control"
	"github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"net/http"
)


func (srv *Service) fillPutFields(c controller.MContext, user *model.User, req *control.PutUserRequest) (fields []string) {
	if len(req.Phone) != 0 {
		if sug := CheckPhone(req.Phone); len(sug) != 0 {
			c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
				Code:  types.CodeBadPhone,
				Error: sug,
			})
			return
		}

		user.Phone = req.Phone
		fields = append(fields, "phone")
	}
	return fields
}
