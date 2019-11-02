package userservice

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	ginhelper "github.com/Myriad-Dreamin/minimum-template/service/gin-helper"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type LoginRequest struct {
	// ID: 用户的id
	ID uint `form:"id" json:"id"`
	// Name: 用户的唯一名称
	Name string `form:"user_name" json:"user_name"`
	// Phone: 用户的邮箱
	Phone string `form:"phone" json:"phone"`
	// Aborted
	//Phone    string `form:"phone" json:"phone"`
	// Password: 用户的密码
	Password string `form:"password" json:"password" xorm:"'password'" binding:"required"`
}

type LoginReply struct {
	Code         int      `json:"code"`
	Identity     []string `json:"identity"`
	Phone        string   `json:"phone"`
	ID           uint     `json:"id"`
	NickName     string   `json:"nick_name"`
	Name         string   `json:"name"`
	Token        string   `json:"token"`
	RefreshToken string   `json:"refresh_token"`
}

func UserToLoginReply(user *model.User, token, refreshToken string, identities []string) *LoginReply {
	return &LoginReply{
		Code:         types.CodeOK,
		Identity:     identities,
		Phone:        user.Phone,
		ID:           user.ID,
		NickName:     user.NickName,
		Name:         user.Name,
		Token:        token,
		RefreshToken: refreshToken,
	}
}

/**
Login v1/user/login POST

- 以下三个属性任选一个，提供多个按照顺序处理
	- `id` uint: the id of logging user
	- `user_name` string: the name of logging user
	- `phone` string: the phone of logging user
- `password` string: the password of logging user

returns:
- `code` int: the operation results
	- types.CodeBindError(1): wrong input, description will be
	attached to the segment of `error`
	- types.CodeNotFound(101): the logging user's information is missing in
	the database
	- types.CodeUserIDMissing(10000): the identity of login user could not
	be detected, one of {id|user_name|phone} must be in the params
	- types.CodeUserWrongPassword(10001): password authenticate failed
- `error` string: options description of bad code
- `user_name` string: the name of logging user
- `identity` list[string]: including all the identities in the system,
	e.g. ["teacher","staff" ... ]
- `phone` string: the phone of logging user
- `id` uint: the id of logging user
- `nick_name` string: the nick name of logging user

Internal Error:
- types.CodeAuthGenerateTokenError(1000): error occurs when
generating the token for logging
- types.CodeAuthenticatePasswordError(1001): error occurs when
authenticating the password
*/

func (srv *Service) Login(c *gin.Context) {
	var req = new(LoginRequest)

	if !ginhelper.BindRequest(c, req) {
		return
	}

	var user *model.User
	var err error
	if req.ID != 0 {
		user, err = srv.db.Query(req.ID)
	} else if len(req.Name) != 0 {
		user, err = srv.db.QueryName(req.Name)
	} else if len(req.Phone) != 0 {
		user, err = srv.db.QueryPhone(req.Phone)
	} else {
		c.JSON(http.StatusOK, &ginhelper.Response{
			Code: types.CodeUserIDMissing,
		})
		return
	}
	if ginhelper.MaybeSelectError(c, user, err) {
		return
	}

	if !ginhelper.AuthenticatePassword(c, user, req.Password) {
		return
	}

	if token, refreshToken, err := srv.middleware.GenerateTokenWithRefreshToken(&types.CustomFields{UID: int64(user.ID)}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ginhelper.ErrorSerializer{
			Code:  types.CodeAuthGenerateTokenError,
			Error: err.Error(),
		})
		return
	} else {
		user.LastLogin = time.Now()

		var identities []string
		if srv.enforcer.HasGroupingPolicy("user:"+strconv.Itoa(int(user.ID)), "admin") {
			identities = append(identities, "admin")
		}

		c.JSON(http.StatusOK, UserToLoginReply(user, token, refreshToken, identities))

		aff, err := user.UpdateFields([]string{"last_login"})
		if err != nil || aff == 0 {
			srv.logger.Debug("update last login failed", "error", err, "affected", aff)
		}

		return
	}
}
