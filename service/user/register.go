package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/model"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"net/http"
	"strconv"
)

type RegisterRequest struct {
	// UserName: 注册用户的名字
	Name string `form:"name" json:"name" binding:"required"`
	// Password: 密码
	Password string `form:"password" json:"password" binding:"required"`
	// NickName: 昵称
	NickName string `form:"nick_name" json:"nick_name" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
	//Email string `form:"email" json:"email" binding:"email"`
	// Gender: 0表示保密, 1表示女, 2表示男, 3~255表示其他
	//Gender uint8 `form:"gender" json:"gender"`
}

type RegisterReply struct {
	// Code: 操作的结果
	Code types.CodeType `json:"code"`
	// ID: 用户的id
	ID uint `json:"id"`
}

func (p RegisterReply) GetID() uint {
	return p.ID
}

func (srv *Service) Register(c controller.MContext) {
	var req = new(RegisterRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	if sug := CheckStrongPassword(req.Password); len(sug) != 0 {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeWeakPassword,
			Error: sug,
		})
		return
	}

	if sug := CheckPhone(req.Phone); len(sug) != 0 {
		c.AbortWithStatusJSON(http.StatusOK, types.ErrorSerializer{
			Code:  types.CodeBadPhone,
			Error: sug,
		})
		return
	}

	
	var user = new(model.User)
	user.Name = req.Name
	user.Password = req.Password

	user.NickName = req.NickName
	user.Phone = req.Phone
	//user.Gender = req.Gender
	//user.Email = req.Email

	// check default value
	aff, err := user.Register()
	if err != nil {
		//fmt.Println(reflect.TypeOf(err))
		if ginhelper.CheckInsertError(c, err) {
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &types.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: err.Error(),
		})
		return
	} else if aff == 0 {
		c.JSON(http.StatusOK, &types.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: "existed",
		})
		return
	}
	c.JSON(http.StatusOK, srv.AfterPost(&RegisterReply{
		Code: types.CodeOK,
		ID:   user.ID,
	}))

	_, err = rbac.AddGroupingPolicy("user:"+strconv.Itoa(int(user.ID)), "norm")
	if err != nil {
		srv.logger.Debug("update group error", "error", err)
	}
}
