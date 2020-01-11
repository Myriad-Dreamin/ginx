package mcore

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

type GormModule struct {
	GormDB *gorm.DB
}

func (m *GormModule) FromRaw(db *gorm.DB, dep module.Module) bool {
	m.GormDB = db
	dep.Provide(DefaultNamespace.DBInstance.GormDB, db)
	return true
}

func (m *GormModule) FromContext(dep module.Module) bool {
	m.GormDB = dep.Require(DefaultNamespace.DBInstance.GormDB).(*gorm.DB)
	return true
}

func (m *GormModule) Install(dep module.Module) bool {
	return m.FromContext(dep)
}

func (m *GormModule) GetGormInstance() *gorm.DB {
	return m.GormDB
}
