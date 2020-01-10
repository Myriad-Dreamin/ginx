package control

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
)

//go:generate package-attach-to-path -generate_register_map

var ObjectCates []interface{}

type ObjectService interface {
	ObjectSignatureXXX() interface{}

	// @Title Do
	// @Description Do Object
	Do(c controller.MContext)
}
