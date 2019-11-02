package userservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/gin-gonic/gin"
)

type PutRequest struct {
	Phone string `form:"phone" json:"phone"`
}

func (srv *Service) fillPutFields(_ *gin.Context, user *model.User, req *PutRequest) (fields []string) {
	if len(req.Phone) != 0 {
		user.Phone = req.Phone
		fields = append(fields, "phone")
	}
	return fields
}
