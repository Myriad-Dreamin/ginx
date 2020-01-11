package mcore

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-template/lib/core-cfg"
)

type DormModule struct {
	DormDB *dorm.DB
}

func (m *DormModule) FromRaw(db *dorm.DB, dep module.Module) bool {
	m.DormDB = db
	dep.Provide(DefaultNamespace.DBInstance.DormDB, m.DormDB)
	return true
}


func (m *DormModule) FromRawSQL(db *sql.DB, dep module.Module) bool {
	logger := dep.Require(DefaultNamespace.Global.Logger).(logger.Logger)
	xdb, err := dorm.FromRaw(db,
		adapt(logger), dorm.Escaper(m.getEscaper(dep)))
	
	m.DormDB = xdb
	dep.Provide(DefaultNamespace.DBInstance.DormDB, xdb)

	return Maybe(dep, "init dorm error", err)
}

func (m *DormModule) FromContext(dep module.Module) bool {
	m.DormDB = dep.Require(DefaultNamespace.DBInstance.DormDB).(*dorm.DB)
	return true
}

func (m *DormModule) Install(dep module.Module) bool {
	return m.FromContext(dep)
}

func (m *DormModule) GetDormInstance() *dorm.DB {
	return m.DormDB
}

type DatabaseConfiguration interface {
	GetDatabaseConfiguration() core_cfg.DatabaseConfig
}

func (m *DormModule) getEscaper(dep module.Module) string {
	return dep.Require(DefaultNamespace.Global.Configuration).
		(DatabaseConfiguration).GetDatabaseConfiguration().Escaper
}

type L struct {
	logger.Logger
}

func (l *L) With(kvs ...interface{}) dorm.Logger {
	return &L{l.Logger.With(kvs)}
}

func adapt(logger logger.Logger) dorm.Logger {
	return &L{logger}
}
