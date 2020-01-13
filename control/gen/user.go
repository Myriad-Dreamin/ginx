package main

import (
	"github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

type UserCategories struct {
	serial.VirtualService
	List           *serial.Category
	Login          *serial.Category
	Register       *serial.Category
	ChangePassword *serial.Category
	Inspect        *serial.Category
	IdGroup        *serial.Category
}

func DescribeUserService(cat *serial.Category) serial.ProposingService {
	var userModel = new(model.User)
	var vUserModel model.User
	svc := &UserCategories{
		List: serial.Ink().
			Path("user-list").
			Method(serial.POST, "ListUsers",
				serial.Request(
					serial.Transfer("ListUsersRequest", model.Filter{}),
				),
				serial.Reply(
					codeField,
					serial.ArrayParam(serial.Param("users", serial.Object(
						"ListUserReply",
						serial.Param("nick_name", &vUserModel.NickName),
						serial.Param("last_login", &vUserModel.LastLogin),
					))),
				),
			),
		Login: serial.Ink().
			Path("login").
			Method(serial.POST, "Login",
				serial.Request(
					serial.Param("id", &userModel.ID),
					serial.Param("nick_name", &userModel.NickName),
					serial.Param("phone", &userModel.Phone),
					serial.Param("password", &serial.String, required),
				),
				serial.Reply(
					codeField,
					serial.Param("id", &userModel.ID),
					serial.Param("identity", &serial.Strings),
					serial.Param("phone", &userModel.Phone),
					serial.Param("nick_name", &userModel.NickName),
					serial.Param("name", &userModel.Name),
					serial.Param("token", &serial.String),
					serial.Param("refresh_token", &serial.String),
				),
			),
		Register: serial.Ink().
			Path("register").
			Method(serial.POST, "Register",
				serial.Request(
					serial.Param("name", &serial.String, required),
					serial.Param("password", &serial.String, required),
					serial.Param("nick_name", &serial.String, required),
					serial.Param("phone", &serial.String, required),
				),
				serial.Reply(
					codeField,
					serial.Param("id", &userModel.ID)),
			),
		ChangePassword: serial.Ink().
			Path("user/:id/password").
			Method(serial.PUT, "ChangePassword",
				serial.Request(
					serial.Param("old_password", &serial.String, required),
					serial.Param("new_password", &serial.String, required),
				),
			),
		Inspect: serial.Ink().Path("user/:id/inspect").
			Method(serial.GET, "InspectUser",
				serial.Reply(
					codeField,
					serial.Param("user", &userModel),
				),
			),
		IdGroup: serial.Ink().
			Path("user/:id").
			Method(serial.GET, "GetUser",
				serial.Reply(
					codeField,
					serial.Param("nick_name", &userModel.NickName),
					serial.Param("last_login", &userModel.LastLogin),
				)).
			Method(serial.PUT, "PutUser",
				serial.Request(
					serial.Param("phone", &userModel.Phone),
				)).
			Method(serial.DELETE, "Delete"),
	}
	svc.Name("UserService").CateOf(cat).UseModel(
		serial.Model(serial.Name("user"), &userModel),
		serial.Model(serial.Name("vUser"), &vUserModel))
	return svc
}
