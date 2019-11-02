package tester

import (
	"encoding/json"
	"errors"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/Myriad-Dreamin/minimum-template/server"
	userservice "github.com/Myriad-Dreamin/minimum-template/service/user"
	"io"
	"log"
	"testing"
)

type Tester struct {
	*server.Mocker
}

type TesterContext struct {
	*server.MockerContext
	t *testing.T
}

func StartTester(serverOptions []server.Option) (tester *Tester) {
	tester = new(Tester)
	tester.Mocker = server.Mock(serverOptions...)
	if tester.Mocker == nil {
		panic(errors.New("req mocker error"))
	}
	return tester
}

func (t *Tester) Context(tt *testing.T) (s *TesterContext) {
	return &TesterContext{
		MockerContext: t.Mocker.Context(tt),
		t:             tt,
	}
}

func (t *TesterContext) AssertNoError(noErr bool) *TesterContext {
	t.MockerContext = t.MockerContext.AssertNoError(noErr)
	return t
}

type ErrorObject struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func (t *TesterContext) DecodeJSON(body io.Reader, req interface{}) interface{} {
	if err := json.NewDecoder(body).Decode(req); err != nil {
		t.t.Fatal(err)
	}
	return req
}

func (t *Tester) Release() {
	t.Mocker.ReleaseMock()
}

func (t *Tester) MakeAdminContext() bool {
	resp := t.Post("/v1/user", userservice.RegisterRequest{
		Name:         "admin_context",
		Password:     "admin",
		NickName:     "admin_context",
		Phone:        "1234567891011",
		RegisterCity: "Qing Dao S.D.",
	})
	if !t.NoErr(resp) {
		return false
	}

	var r userservice.RegisterReply
	err := resp.JSON(&r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	resp = t.Post("/v1/login",
		userservice.LoginRequest{
			ID:       r.ID,
			Password: "admin",
		}, mock.Comment("admin login for test"))
	if !t.NoErr(resp) {
		return false
	}

	var r2 userservice.LoginReply
	err = resp.JSON(&r2)
	if err != nil {
		log.Fatal(err)
		return false
	}

	//fmt.Println(r2)
	//r2.RefreshToken
	t.UseToken(r2.Token)
	return true
}

func (t *Tester) MainM(m *testing.M) {
	t.Main(func() {
		m.Run()
	})
}

func (t *Tester) Main(doSomething func()) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			t.Logger.Error("panic", "error", err)
		}
		t.Release()
	}()
	if !t.MakeAdminContext() {
		return
	}
	doSomething()
}
