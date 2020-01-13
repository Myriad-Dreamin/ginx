package main

import (
	"fmt"
	serial "github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

var codeField = serial.Param("code", new(types.CodeRawType))
var required = serial.Tag("binding", "required")

func main() {
	V1Cate := serial.Ink().Path("v1")

	userCate := DescribeUserService(V1Cate)
	userCate.ToFile("user.go")
	objectCate := DescribeObjectService(V1Cate)
	objectCate.ToFile("object.go")
	err := serial.NewService(
		userCate, objectCate).Publish()
	if err != nil {
		fmt.Println(err)
	}
}

