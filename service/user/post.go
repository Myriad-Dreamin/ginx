package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control/auth"
	base_service "github.com/Myriad-Dreamin/minimum-template/lib/base-service"
	"github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/model"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

type PostReply struct {
	Code types.CodeRawType `json:"code"`
	User *model.User    `json:"user"`
}

func (p PostReply) GetID() uint {
	return p.User.ID
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

func (srv *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req PostRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.User)
	// fill here
	return obj
}

type PostReplyI interface {
	GetID() uint
}

func (srv *Service) AfterPost(reply PostReplyI) interface{} {
	if b, err := auth.UserEntity.AddReadPolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.GetID()), reply.GetID()); err != nil {
		if !b {
			srv.logger.Debug("add failed")
		}
		return serial.ErrorSerializer{
			Code:  types.CodeAddReadPrivilegeError,
			Error: err.Error(),
		}
	} else {
		if !b {
			srv.logger.Debug("add failed")
		}
	}

	if b, err := auth.UserEntity.AddWritePolicy(srv.enforcer, auth.UserEntity.CreateObj(reply.GetID()), reply.GetID()); err != nil {
		if !b {
			srv.logger.Debug("add failed")
		}
		return serial.ErrorSerializer{
			Code:  types.CodeAddWritePrivilegeError,
			Error: err.Error(),
		}
	} else {
		if !b {
			srv.logger.Debug("add failed")
		}
	}
	return reply
}
