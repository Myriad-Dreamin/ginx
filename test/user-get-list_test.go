package tests

import (
	"fmt"
	userservice "github.com/Myriad-Dreamin/minimum-template/service/user"
	"testing"
)

func testUserList(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)

	resp := srv.Get("/v1/user-list?page=1&page_size=1")

	reply := srv.DecodeJSON(resp.Body(), new(userservice.ListReply)).(*userservice.ListReply)
	fmt.Println(reply)
}
