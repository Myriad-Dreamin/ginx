package tests

import (
	"fmt"
	userservice "github.com/Myriad-Dreamin/minimum-template/service/user"
	"testing"
)

func testGet(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	resp := srv.Get("/v1/user/1")
	reply := srv.DecodeJSON(resp.Body(), new(userservice.GetReply)).(*userservice.GetReply)
	fmt.Println(reply, reply.User)
}
