package model

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-template/model/sp-layer"
	
)

func InstallFromContext(dep module.Module) bool {
	return dblayer.FromContext(dep)
}

func Install(dep module.Module) bool {
	return dblayer.Install(dep)
}

func InstallMock(dep module.Module) bool {
	return dblayer.InstallMock(dep)
}

func RegisterRedis(dep module.Module) bool {
	return splayer.RegisterRedis(dep)
}

func InstallRedis(dep module.Module) bool {
	return splayer.Install(dep)
}

func Close(dep module.Module) bool {
	x := dblayer.Close(dep)
	x = x && splayer.Close(dep)
	return x
}


type Provider = splayer.Provider

func NewProvider(namespace string) *Provider {
	return splayer.NewProvider(namespace)
}

func SetProvider(p *Provider) *Provider {
	return splayer.SetProvider(p)
}

