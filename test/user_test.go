package tests

import (
	"github.com/Myriad-Dreamin/minimum-template/test/tester"
	"testing"
)

func TestUser(t *testing.T) {
	_ = t.Run("RegisterLogin", srv.HandleTestWithOutError(testUserRegisterLogin)) &&
		t.Run("Get", testUserGet) &&
		t.Run("Inspect", srv.HandleTestWithOutError(testUserInspect)) &&
		t.Run("List", testUserList) &&
		t.Run("ChangePassword", srv.HandleTestWithOutError(testUserChangePassword)) &&
		t.Run("Put", srv.HandleTestWithOutError(testUserPut)) &&
		t.Run("Delete", srv.HandleTestWithOutError(testUserDelete))
		
}

func testUserInspect(t *tester.TesterContext) {

}

func testUserPut(t *tester.TesterContext) {

}
