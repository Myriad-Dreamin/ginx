package main

import (
	"github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

type ObjectCategories struct {
	serial.VirtualService
	List    serial.Category
	Post    serial.Category
	Inspect serial.Category
	IdGroup serial.Category
}

func DescribeObjectService(base string) serial.ProposingService {
	var objectModel = new(model.Object)
	svc := &ObjectCategories{
		List: serial.Ink().
			Path("object-list").
			Method(serial.POST, "ListObjects",
				serial.Request(
					serial.Transfer("ListObjectsRequest", model.Filter{}),
				),
				serial.Reply(
					codeField,
					serial.ArrayParam(serial.Param("objects", objectModel)),
				),
			),
		Post: serial.Ink().
			Path("object").
			Method(serial.POST, "PostObject",
				serial.Request(),
				serial.Reply(
					codeField,
					serial.Param("object", &objectModel),
				),
			),
		Inspect: serial.Ink().Path("object/:oid/inspect").
			Method(serial.GET, "InspectObject",
				serial.Reply(
					codeField,
					serial.Param("user", &objectModel),
				),
			),
		IdGroup: serial.Ink().
			Path("object/:oid").
			Method(serial.GET, "GetObject",
				serial.Reply(
					codeField,
					serial.Param("object", &objectModel),
				)).
			Method(serial.PUT, "PutObject",
				serial.Request()).
			Method(serial.DELETE, "Delete"),
	}
	svc.Name("ObjectService").Base(base) //.
	//UseModel(serial.Model(serial.Name("object"), &objectModel))
	return svc
}
