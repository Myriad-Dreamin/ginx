package userservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	base_service "github.com/Myriad-Dreamin/ginx/service/base-service"
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
)

type PostReply struct {
	Code int         `json:"code"`
	User *model.User `json:"user"`
}

func UserToPostReply(obj *model.User) *PostReply {
	return &PostReply{
		Code: types.CodeOK,
		User: obj,
	}
}

/*
1）	创建用户类：用户标识、用户名、登录密码、用户类型（系统管理员/普通用户）、用户姓名、手机号码(11位数字)、用户级别（钻石级、重要、一般）、注册城市（统计分析时可以考虑）、注册时间、修改时间。在后台数据库至少要建立一个管理员用户,用户名：admin,密码admin。
*/

type PostRequest struct {
}

func (srv *Service) SerializePost(c *gin.Context) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.User)
	// fill here
	return obj
}
