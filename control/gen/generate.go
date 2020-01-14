package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-template/lib/artist"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

var codeField = artist.Param("code", new(types.CodeRawType))
var required = artist.Tag("binding", "required")

func main() {
	v1 := "v1"

	userCate := DescribeUserService(v1)
	userCate.ToFile("user.go")
	objectCate := DescribeObjectService(v1)
	objectCate.ToFile("object.go")
	err := artist.NewService(
		userCate, objectCate).Publish()
	if err != nil {
		fmt.Println(err)
	}
}
