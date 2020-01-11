package mcore

import (
	"database/sql"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type RawSQLModule struct {
	RawDB *sql.DB
}

func (m *RawSQLModule) FromRaw(rawDB *sql.DB, dep module.Module) bool {
	m.RawDB = rawDB
	dep.Provide(DefaultNamespace.DBInstance.RawDB, rawDB)
	return true
}

func (m *RawSQLModule) GetRawSQLInstance() *sql.DB {
	return m.RawDB
}
