package tests

import (
	userservice "github.com/Myriad-Dreamin/minimum-template/service/user"
	"github.com/Myriad-Dreamin/minimum-template/test/tester"
)

func testUserRegisterLogin(t *tester.TesterContext) {
	var (
		name  = "chan tan"
		nick  = "tan chan"
		phone = "10086111"
		pswd  = normalUserPassword
	)
	resp := t.Post("/v1/user", userservice.RegisterRequest{
		Name:     name,
		Password: pswd,
		NickName: nick,
		Phone:    phone,
	})
	id := t.DecodeJSON(resp.Body(),
		new(userservice.RegisterReply)).(*userservice.RegisterReply).ID
	resp = t.Post("/v1/login", userservice.LoginRequest{
		ID:       id,
		Password: pswd,
	})
	resp = t.Post("/v1/login", userservice.LoginRequest{
		NickName: nick,
		Password: pswd,
	})
	resp = t.Post("/v1/login", userservice.LoginRequest{
		Phone:    phone,
		Password: pswd,
	})

	srv.Set(normalUserIdKey, id)
}
