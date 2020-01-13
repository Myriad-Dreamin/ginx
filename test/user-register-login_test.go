package tests

import (
	"github.com/Myriad-Dreamin/minimum-template/control"
	"github.com/Myriad-Dreamin/minimum-template/test/tester"
)

func testUserRegisterLogin(t *tester.TesterContext) {
	var (
		name  = "chan tan"
		nick  = "tan chan"
		phone = "10086111"
		pswd  = normalUserPassword
	)
	resp := t.Post("/v1/user", control.RegisterRequest{
		Name:     name,
		Password: pswd,
		NickName: nick,
		Phone:    phone,
	})
	id := t.DecodeJSON(resp.Body(),
		new(control.RegisterReply)).(*control.RegisterReply).Id
	resp = t.Post("/v1/login", control.LoginRequest{
		Id:       id,
		Password: pswd,
	})
	resp = t.Post("/v1/login", control.LoginRequest{
		NickName: nick,
		Password: pswd,
	})
	resp = t.Post("/v1/login", control.LoginRequest{
		Phone:    phone,
		Password: pswd,
	})

	srv.Set(normalUserIdKey, id)
}
