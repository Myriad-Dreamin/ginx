package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

// UserService defines the interface of user service
// @Category User - Login
// @Description Login with password
// @Path /v1/login
type userLoginCate interface {
}

// @Category User - Register
// @Description Register
// @Path /v1/user
type userRegisterCate interface {
}

// @Category User - Change Password
// @Description Change Password
// @Path /v1/user/:id/password
type userCgPasswordCate interface {
}

// @Category User - Delete/Get Api Group
// @Description hint: the Delete api is Admin-only callable
// @Path /v1/user/:id
type userIdGroupCate interface {
}

// @Category User - List
// @Path /v1/user-list
type userListCate interface {
}

//// @Category User - Change Password
//// @Description Change Password
//// @Path /v1/user/:id/password
//type userCate interface {
//}

var UserCates []interface{}

func init() {
	var (
		loginCate    userLoginCate      = 0
		registerCate userRegisterCate   = 0
		cgCate       userCgPasswordCate = 0
		d            userCgPasswordCate = 0
		i            userIdGroupCate    = 0
		j            userListCate       = 0
	)
	UserCates = []interface{}{
		&loginCate,
		&registerCate,
		&cgCate,
		&d, &i, &j, //&k,
	}
}

type UserService interface {
	UserSignatureXXX() interface{}

	// @Title Login
	// @Description Login to access more api in backend.
	// The following is a description of the parameters
	//     + `password string`: password of the user
	// uid, one of following three parameter must be provided:
	//         + `id uint`: id of the user
	//
	//         + `nick_name string`: nick-name of the user
	//
	//         + `phone string` phone number of the user
	//
	// The following is a description of the returns
	// + `identity array[string]`: the groups user currently in
	// + `id uint`: the unique user id in database
	// + `phone string`: the phone of user
	// + `nick_name string`: the unique user nick name in this website
	// + `name string`: the true name of user
	// + `token string`: the jwt token identifies an user, which must set in every auth api request's header
	// + `refresh_token string`: the jwt token used to refresh token without requirement of user inputting password to login again
	Login(c controller.MContext)

	// @Title Register
	// @Description Register an user in minimum-template server
	// The following is a description of the parameters
	// + `name string`: name of the user
	// + `nick_name string`: nick-name of the user, must be unique
	// + `phone string` phone number of the user, must be unique
	// + `password string` password number of the user, which must pass the [password test](https://github.com/Myriad-Dreamin/minimum-template/blob/master/service/user/change-password.go).
	// + `city_code string`: code of register city of the user
	//
	// The following is a description of the returns
	//     + `id uint`: the unique user id in database
	Register(c controller.MContext)

	// @Title Put
	// @Description Update User Information. only the phone number can be modified
	// The following is a description of the parameters
	//     + `phone string` phone number of the user
	Put(c controller.MContext)

	// @Title ChangePassword
	// @Description change password of user
	// The following is a description of the parameters
	//     + `old_password string` old password number of the user, must match the string in database after encrypted
	//     + `new_password string` new password number of the user, which must pass the [password test](https://github.com/Myriad-Dreamin/minimum-template/blob/master/service/user/change-password.go).
	ChangePassword(c controller.MContext)

	//// /:id/grant PUT
	//Grant(c mcontext.MContext)

	// @Title Get
	// @Description Get User
	Get(c controller.MContext)

	// @Title Delete
	// @Description Delete User
	Delete(c controller.MContext)

	// @Title List
	// @Description List User
	List(c controller.MContext)
}
