package tests

import (
	"github.com/Myriad-Dreamin/minimum-template/model"
	userservice "github.com/Myriad-Dreamin/minimum-template/service/user"
	"testing"
)

func testRegisterLogin(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	var (
		name  = "chan tan"
		nick  = "tan chan"
		phone = "10086111"
		pswd  = "11122222"
	)
	resp := srv.Post("/v1/user", userservice.RegisterRequest{
		Name:         name,
		Password:     pswd,
		NickName:     nick,
		Phone:        phone,
		RegisterCity: "tan arch",
	})
	id := srv.DecodeJSON(resp.Body(),
		new(userservice.RegisterReply)).(*userservice.RegisterReply).ID
	resp = srv.Post("/v1/login", userservice.LoginRequest{
		ID:       id,
		Password: pswd,
	})
	resp = srv.Post("/v1/login", userservice.LoginRequest{
		Name:     name,
		Password: pswd,
	})
	resp = srv.Post("/v1/login", userservice.LoginRequest{
		Phone:    phone,
		Password: pswd,
	})
	_, err := (&model.User{ID: id}).Delete()
	if !srv.NoError(err, "delete error") {
		return
	}
}
