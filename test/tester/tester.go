package tester

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/Myriad-Dreamin/minimum-template/control"
	"github.com/Myriad-Dreamin/minimum-template/server"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"io"
	"log"
	"strconv"
	"testing"
)

type Tester struct {
	*server.Mocker

	ContextVars map[string]interface{}
}

type TesterContext struct {
	*server.MockerContext
	t *testing.T
	sugar.HandlerErrorLogger
}

func (tester Tester) Set(k string, v interface{}) interface{} {
	res, _ := tester.ContextVars[k]
	tester.ContextVars[k] = v
	return res
}

func (tester Tester) Get(k string) interface{} {
	return tester.ContextVars[k]
}

func (tester Tester) ShouldGet(k string) (v interface{}, ok bool) {
	v, ok = tester.ContextVars[k]
	return
}

func (tester Tester) MustGet(k string) interface{} {
	v, ok := tester.ContextVars[k]
	if !ok {
		panic(fmt.Errorf("could not get %v from context", k))
	}
	return v
}

func StartTester(serverOptions []server.Option) (tester *Tester) {
	tester = new(Tester)
	tester.ContextVars = make(map[string]interface{})
	tester.Mocker = server.Mock(serverOptions...)
	if tester.Mocker == nil {
		panic(errors.New("req mocker error"))
	}
	return tester
}

func (tester *Tester) Context(tt *testing.T) (s *TesterContext) {
	return &TesterContext{
		MockerContext:      tester.Mocker.Context(tt),
		t:                  tt,
		HandlerErrorLogger: sugar.NewHandlerErrorLogger(tt),
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

func (tester *Tester) Release() {
	tester.Mocker.ReleaseMock()
}

func (tester *Tester) MakeAdminContext() bool {
	resp := tester.Post("/v1/user", control.RegisterRequest{
		Name:     "admin_context",
		Password: "Admin12345678",
		NickName: "admin_context",
		Phone:    "1234567891011",
	}, mock.Comment("admin register for test"))
	if !tester.NoErr(resp) {
		return false
	}

	var r control.RegisterReply
	err := resp.JSON(&r)
	if err != nil {
		log.Fatal(err)
		return false
	}
	resp = tester.Post("/v1/login",
		control.LoginRequest{
			Id:       r.Id,
			Password: "Admin12345678",
		}, mock.Comment("admin login for test"))
	if !tester.NoErr(resp) {
		return false
	}

	var r2 control.LoginReply
	err = resp.JSON(&r2)
	if err != nil {
		log.Fatal(err)
		return false
	}

	//fmt.Println(r2)
	//r2.RefreshToken
	_, err = rbac.AddGroupingPolicy("user:"+strconv.Itoa(int(r2.Id)), types.GroupAdmin)
	if err != nil {
		tester.Logger.Debug("update group error", "error", err)
	}
	//fmt.Println("QAQQQ", rbac.GetPolicy())
	//fmt.Println("QAQQQ", rbac.GetGroupingPolicy())
	tester.UseToken(r2.Token)
	return true
}

func (tester *Tester) MainM(m *testing.M) {
	tester.Main(func() {
		m.Run()
	})
}

func (tester *Tester) Main(doSomething func()) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			tester.Logger.Error("panic", "error", err)
		}
		tester.Release()
	}()
	if !tester.MakeAdminContext() {
		return
	}
	doSomething()
}

type GoStyleTestFunc func(*testing.T)
type MinimumStyleTestFunc func(ctx *TesterContext)

func (tester *Tester) HandleTest(testFunc MinimumStyleTestFunc) GoStyleTestFunc {
	return func(t *testing.T) {
		testFunc(tester.Context(t))
	}
}

func (tester *Tester) HandleTestWithOutError(testFunc MinimumStyleTestFunc) GoStyleTestFunc {
	return func(t *testing.T) {
		testFunc(tester.Context(t).AssertNoError(true))
	}
}
