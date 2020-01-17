package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

var codeField = artisan.Param("code", new(types.CodeRawType))
var required = artisan.Tag("binding", "required")

func main() {
	v1 := "v1"

	//instantiate
	userCate := DescribeUserService(v1)
	objectCate := DescribeObjectService(v1)

	//to files
	userCate.ToFile("user.go")
	objectCate.ToFile("object.go")

	err := artisan.NewService(
		userCate,
		objectCate,
	).Publish()
	if err != nil {
		fmt.Println(err)
	}
}
