package mcore


import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
)


type ModelCallbacks interface {
	Migrates() error
	Injects() error
}

func ModelCallback(callback ModelCallbacks, dep module.Module) bool {
	return true &&
		Maybe(dep, "inject callback error", callback.Injects()) &&
		Maybe(dep, "migrate callback error", callback.Migrates())
}
