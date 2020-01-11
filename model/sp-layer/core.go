package splayer

import (
	"github.com/Myriad-Dreamin/minimum-template/lib/core"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

var p = newModelModule()

func RegisterRedis(dep module.Module) bool {
	return p.Install(dep)
}

func FromContext(dep module.Module) bool {
	return p.FromContext(dep)
}

func Install(dep module.Module) bool {
	return p.Install(dep)
}

func Close(dep module.Module) bool {
	if p.Opened {
		if err := p.RedisPool.Close(); err != nil {
			p.Logger.Error("close Redis error", "error", err)
			return false
		}
	}
	return true
}

type modelModule struct {
	mcore.RedisPoolModule
	mcore.LoggerModule
	Opened bool
}

func newModelModule() modelModule {
	return modelModule{
		Opened: false,
	}
}

func (m *modelModule) FromContext(dep module.Module) bool {
	m.Opened = true && 
		m.LoggerModule.Install(dep) &&
		m.RedisPoolModule.FromContext(dep)
	return m.Opened
}

func (m *modelModule) Install(dep module.Module) bool {
	m.Opened = true && 
		m.LoggerModule.Install(dep) &&
		m.RedisPoolModule.InstallFromConfiguration(dep)
	return m.Opened
}
