package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-template/lib/artisan"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

var codeField = artisan.Param("code", new(types.CodeRawType))
var required = artisan.Tag("binding", "required")

func main() {
	v1 := "v1"

	userCate := DescribeUserService(v1)
	userCate.ToFile("user.go")
	objectCate := DescribeObjectService(v1)
	objectCate.ToFile("object.go")
	err := artisan.NewService(
		userCate, objectCate).Publish()
	if err != nil {
		fmt.Println(err)
	}
}
