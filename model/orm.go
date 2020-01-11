package model

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	
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

func Close(dep module.Module) bool {
	return dblayer.Close(dep)
}

