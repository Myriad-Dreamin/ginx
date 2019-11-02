package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"github.com/Myriad-Dreamin/ginx/model"
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RegisterRequest struct {
	// UserName: 注册用户的名字
	Name string `form:"name" json:"name" binding:"required"`
	// Password: 密码
	Password string `form:"password" json:"password" binding:"required"`
	// NickName: 昵称
	NickName     string `form:"nick_name" json:"nick_name" binding:"required"`
	Phone        string `form:"phone" json:"phone" binding:"required"`
	RegisterCity string `form:"register_city" json:"register_city" binding:"required"`
	//Email string `form:"email" json:"email" binding:"email"`
	// Gender: 0表示保密, 1表示女, 2表示男, 3~255表示其他
	//Gender uint8 `form:"gender" json:"gender"`
}

type RegisterReply struct {
	// Code: 操作的结果
	Code int `json:"code"`
	// ID: 用户的id
	ID uint `json:"id"`
}

/**
Register v1/user/register POST

params:
- `name` string: the name of registering user
- `password` string: the password of this user's account
- `nick_name` string: the nick_name of this user
- `gender` uint8: 0 and 3~255 represents for secret,1 represents for female,2 represents
	for male

returns:
- `code` int: the operation results
	- types.CodeBindError(1): wrong input, description will be
	attached to the segment of `error`
	- types.CodeInsertError(100): insert error, description will be
	attached to the segment of `error`
	- types.CodeDuplicatePrimaryKey(104): insert error, throw by mysql plugin
- `error` string: options description of bad code
- `id` uint: the id of this user
*/
func (srv *Service) Register(c *gin.Context) {
	var req = new(RegisterRequest)
	if !ginhelper.BindRequest(c, req) {
		return
	}

	var user = new(model.User)
	user.Name = req.Name
	/**
	要求密码不少于6位，必须含有两个数字，不能与原来的密码相似，不能都为大写或小写
	*/
	user.Password = req.Password

	user.NickName = req.NickName
	user.Phone = req.Phone
	user.RegisterCity = req.RegisterCity
	//user.Gender = req.Gender
	//user.Email = req.Email

	// check default value
	aff, err := user.Register()
	if err != nil {
		//fmt.Println(reflect.TypeOf(err))
		if ginhelper.CheckInsertError(c, err) {
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ginhelper.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: err.Error(),
		})
		return
	} else if aff == 0 {
		c.JSON(http.StatusOK, &ginhelper.ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: "existed",
		})
		return
	}
	c.JSON(http.StatusOK, &RegisterReply{
		Code: types.CodeOK,
		ID:   user.ID,
	})

	_, err = rbac.AddGroupingPolicy("user:"+strconv.Itoa(int(user.ID)), "norm")
	if err != nil {
		srv.logger.Debug("update group error", "error", err)
	}
}
