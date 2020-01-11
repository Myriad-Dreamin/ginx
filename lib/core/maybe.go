package mcore

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
)


type ErrorLogger interface {
	Error(msg string, keyvals ...interface{})
}

func Maybe(dep module.Module, hint string, err error) bool {
	if err != nil {
		logger := dep.Require(DefaultNamespace.Global.Logger).(ErrorLogger)
		logger.Error(hint, "error", err)
		return false
	}
	return true
}
