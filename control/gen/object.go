package main

import (
	"github.com/Myriad-Dreamin/minimum-template/lib/artist"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

type ObjectCategories struct {
	artist.VirtualService
	List    artist.Category
	Post    artist.Category
	Inspect artist.Category
	IdGroup artist.Category
}

func DescribeObjectService(base string) artist.ProposingService {
	var objectModel = new(model.Object)
	svc := &ObjectCategories{
		List: artist.Ink().
			Path("object-list").
			Method(artist.POST, "ListObjects",
				artist.QT("ListObjectsRequest", model.Filter{}),
				artist.Reply(
					codeField,
					artist.ArrayParam(artist.Param("objects", objectModel)),
				),
			),
		Post: artist.Ink().
			Path("object").
			Method(artist.POST, "PostObject",
				artist.Request(),
				artist.Reply(
					codeField,
					artist.Param("object", &objectModel),
				),
			),
		Inspect: artist.Ink().Path("object/:oid/inspect").
			Method(artist.GET, "InspectObject",
				artist.Reply(
					codeField,
					artist.Param("object", &objectModel),
				),
			),
		IdGroup: artist.Ink().
			Path("object/:oid").
			Method(artist.GET, "GetObject",
				artist.Reply(
					codeField,
					artist.Param("object", &objectModel),
				)).
			Method(artist.PUT, "PutObject",
				artist.Request()).
			Method(artist.DELETE, "Delete"),
	}
	svc.Name("ObjectService").Base(base) //.
	//UseModel(serial.Model(serial.Name("object"), &objectModel))
	return svc
}
