package main

import (
	"fmt"
	serial "github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

var codeField = serial.Param("code", new(types.CodeRawType))
var required = serial.Tag("binding", "required")

func main() {
	v1 := "v1"

	userCate := DescribeUserService(v1)
	userCate.ToFile("user.go")
	objectCate := DescribeObjectService(v1)
	objectCate.ToFile("object.go")
	err := serial.NewService(
		userCate, objectCate).Publish()
	if err != nil {
		fmt.Println(err)
	}
}
