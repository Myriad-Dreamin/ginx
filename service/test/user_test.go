package tests

import (
	"testing"
)

func TestUser(t *testing.T) {
	_ =
		t.Run("TestRegisterLogin", testRegisterLogin) &&
			t.Run("TestGet", testGet) &&
			t.Run("TestList", testList)
}
