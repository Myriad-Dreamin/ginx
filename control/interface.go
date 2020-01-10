package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

//go:generate package-attach-to-path -generate_register_map

var ObjectCates []interface{}

type ObjectService interface {
	ObjectSignatureXXX() interface{}

	// @Title Post
	// @Description Post Object
	Post(c controller.MContext)

	// @Title Put
	// @Description Put Object
	Put(c controller.MContext)

	// @Title Delete
	// @Description Delete Object
	Delete(c controller.MContext)

	// @Title Get
	// @Description Get Object
	Get(c controller.MContext)

	// @Title List
	// @Description List Object
	List(c controller.MContext)
}
