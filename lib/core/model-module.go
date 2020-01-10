package mcore

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/Myriad-Dreamin/minimum-template/config"
)

type ModelCallbacks interface {
	Migrates() error
	Injects() error	
}

type ModelModule struct {
	DB *gorm.DB
	RawDB *sql.DB
	DormDB *dorm.DB

	cb ModelCallbacks
}

func NewModelModule(cb ModelCallbacks) ModelModule {
	return ModelModule{cb: cb}
}

func (m *ModelModule) GetInstance() *gorm.DB {
	return m.DB
}

func (m *ModelModule) GetRawInstance() *sql.DB {
	return m.RawDB
}

func (m *ModelModule) GetDormInstance() *dorm.DB {
	return m.DormDB
}

func (m *ModelModule) Register(dep module.Module) error {
	var err error
	m.DB = dep.Require(config.ModulePath.Global.GormDB).(*gorm.DB)
	m.RawDB = (*m.DB).DB()

	// test if it is alive
	if err = m.RawDB.Ping(); err != nil {
		return err
	}

	xdb, err := dorm.FromRaw(m.RawDB,
		adapt(dep.Require(config.ModulePath.Global.Logger).(types.Logger)),
		dorm.Escaper(
			dep.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig).DatabaseConfig.Escaper))
	if err != nil {
		return err
	}

	m.DormDB = xdb

	if err := m.cb.Injects(); err != nil {
		return err
	}

	return m.cb.Migrates()
}



