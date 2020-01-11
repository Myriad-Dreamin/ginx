package mcore

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
)

type LoggerModule struct {
	logger logger.Logger
}

func (m *LoggerModule) FromRaw(mlogger logger.Logger, dep module.Module) bool {
	m.logger = mlogger
	dep.Provide(DefaultNamespace.Global.Logger, mlogger)
	return true
}

func (m *LoggerModule) FromContext(dep module.Module) bool {
	m.logger = dep.Require(DefaultNamespace.Global.Logger).(logger.Logger)

	return true
}

func (m *LoggerModule) Install(dep module.Module) bool {
	return m.FromContext(dep)
}

func (m *LoggerModule) GetLoggerInstance() logger.Logger {
	return m.logger
}
