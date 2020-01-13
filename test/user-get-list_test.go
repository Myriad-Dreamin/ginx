package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-template/control"
	"testing"
)

func testUserList(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)

	resp := srv.Get("/v1/user-list?page=1&page_size=1")

	reply := srv.DecodeJSON(resp.Body(), new(control.ListUsersReply)).(*control.ListUsersReply)
	fmt.Println(reply)
}
