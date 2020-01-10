package tests

import (
	userservice "github.com/Myriad-Dreamin/minimum-template/service/user"
	"github.com/Myriad-Dreamin/minimum-template/test/tester"
	"reflect"
	"strconv"
)

func testUserChangePassword(t *tester.TesterContext) {
	id := reflect.ValueOf(srv.Get(normalUserIdKey)).Convert(intType).Interface().(int)
	t.Put("/v1/user/"+strconv.Itoa(id)+"/password", userservice.ChangePasswordRequest{
		OldPassword: normalUserPassword,
		NewPassword: normalUserNewPassword,
	})
}
