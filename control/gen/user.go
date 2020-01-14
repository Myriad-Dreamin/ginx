package main

import (
	"github.com/Myriad-Dreamin/minimum-template/lib/artist"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

type UserCategories struct {
	artist.VirtualService
	List           artist.Category
	Login          artist.Category
	Register       artist.Category
	ChangePassword artist.Category
	Inspect        artist.Category
	IdGroup        artist.Category
}

func DescribeUserService(base string) artist.ProposingService {
	var userModel = new(model.User)
	var vUserModel model.User
	svc := &UserCategories{
		List: artist.Ink().
			Path("user-list").
			Method(artist.POST, "ListUsers",
				artist.QT("ListUsersRequest", model.Filter{}),
				artist.Reply(
					codeField,
					artist.ArrayParam(artist.Param("users", artist.Object(
						"ListUserReply",
						artist.SPsC(&vUserModel.NickName, &vUserModel.LastLogin),
					))),
				),
			),
		Login: artist.Ink().
			Path("login").
			Method(artist.POST, "Login",
				artist.Request(
					artist.SPsC(&userModel.ID, &userModel.NickName, &userModel.Phone),
					artist.Param("password", artist.String, required),
				),
				artist.Reply(
					codeField,
					artist.SPsC(&userModel.ID, &userModel.Phone, &userModel.NickName, &userModel.Name),
					artist.Param("identity", artist.Strings),
					artist.Param("token", artist.String),
					artist.Param("refresh_token", artist.String),
				),
			),
		Register: artist.Ink().
			Path("register").
			Method(artist.POST, "Register",
				artist.Request(
					artist.SPs(artist.C(&userModel.Name, &userModel.NickName, &userModel.Phone), required),
					artist.Param("password", artist.String, required),
				),
				artist.Reply(
					codeField,
					artist.Param("id", &userModel.ID)),
			),
		ChangePassword: artist.Ink().
			Path("user/:id/password").
			Method(artist.PUT, "ChangePassword",
				artist.Request(
					artist.Param("old_password", artist.String, required),
					artist.Param("new_password", artist.String, required),
				),
			),
		Inspect: artist.Ink().Path("user/:id/inspect").
			Method(artist.GET, "InspectUser",
				artist.Reply(
					codeField,
					artist.Param("user", &userModel),
				),
			),
		IdGroup: artist.Ink().
			Path("user/:id").
			Method(artist.GET, "GetUser",
				artist.Reply(
					codeField,
					artist.SPsC(&userModel.NickName, &userModel.LastLogin),
				)).
			Method(artist.PUT, "PutUser",
				artist.Request(
					artist.Param("phone", &userModel.Phone),
				)).
			Method(artist.DELETE, "Delete"),
	}
	svc.Name("UserService").Base(base).UseModel(
		artist.Model(artist.Name("user"), &userModel),
		artist.Model(artist.Name("vUser"), &vUserModel))
	return svc
}
