package splayer

import (
	"github.com/Myriad-Dreamin/minimum-template/lib/core"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

var p = newModelModule()

func RegisterRedis(dep module.Module) bool {
	return p.Install(dep)
}

type modelModule struct {
	mcore.RedisPoolModule
	mcore.LoggerModule
}

func newModelModule() modelModule {
	return modelModule{}
}

func (m *modelModule) Install(dep module.Module) bool {
	return true && 
		m.LoggerModule.Install(dep) &&
		m.RedisPoolModule.FromContext(dep)
}
