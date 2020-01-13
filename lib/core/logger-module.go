package mcore

import (
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type LoggerModule struct {
	Logger logger.Logger
}

func (m *LoggerModule) FromRaw(mlogger logger.Logger, dep module.Module) bool {
	m.Logger = mlogger
	dep.Provide(DefaultNamespace.Global.Logger, mlogger)
	return true
}

func (m *LoggerModule) FromContext(dep module.Module) bool {
	m.Logger = dep.Require(DefaultNamespace.Global.Logger).(logger.Logger)

	return true
}

func (m *LoggerModule) Install(dep module.Module) bool {
	return m.FromContext(dep)
}

func (m *LoggerModule) GetLoggerInstance() logger.Logger {
	return m.Logger
}
