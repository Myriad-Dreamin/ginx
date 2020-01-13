package mcore

import (
	"errors"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-template/lib/core-cfg"
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

func (m *GormModule) installFromConfiguration(
	initFunc func(dep module.Module) (*gorm.DB, error), dep module.Module) bool {
	xdb, err := initFunc(dep)
	m.FromRaw(xdb, dep)
	return Maybe(dep, "init gorm error", err)
}

func (m *GormModule) InstallFromConfiguration(dep module.Module) bool {
	return m.installFromConfiguration(OpenGORM, dep)
}

func (m *GormModule) InstallMockFromConfiguration(dep module.Module) bool {
	return m.installFromConfiguration(MockGORM, dep)
}

func (m *GormModule) GetGormInstance() *gorm.DB {
	return m.GormDB
}

func booleanString(b bool) string {
	if b {
		return "True"
	} else {
		return "False"
	}
}

func concatQueryString(options string) string {
	if len(options) != 0 {
		return "&"
	} else {
		return "?"
	}
}

func getDatabaseConfiguration(dep module.Module) core_cfg.DatabaseConfig {
	return dep.Require(DefaultNamespace.Global.Configuration).(DatabaseConfiguration).GetDatabaseConfiguration()
}

func getRedisConfiguration(dep module.Module) core_cfg.RedisConfig {
	return dep.Require(DefaultNamespace.Global.Configuration).(RedisConfiguration).GetRedisConfiguration()
}

func parseConfig(dep module.Module) (string, string, error) {
	// user:password@/dbname?charset=utf8&parseTime=True&loc=Local

	cfg := getDatabaseConfiguration(dep)

	if len(cfg.ConnectionType) == 0 || len(cfg.User) == 0 || len(cfg.Password) == 0 || len(cfg.DatabaseName) == 0 {
		return "", "", errors.New("not enough params")
	}
	url := cfg.User + ":" + cfg.Password + "@"
	if len(cfg.Host) != 0 {
		url += "(" + cfg.Host + ")"
	}
	url += "/" + cfg.DatabaseName
	options := ""

	if len(cfg.Charset) != 0 {
		options += concatQueryString(options) + "charset=" + cfg.Charset
	}
	if cfg.ParseTime {
		options += concatQueryString(options) + "parseTime=" + booleanString(cfg.ParseTime)
	}
	if len(cfg.Location) != 0 {
		options += concatQueryString(options) + "loc=" + cfg.Location
	}
	return cfg.ConnectionType, url + options, nil
}

func OpenGORM(dep module.Module) (*gorm.DB, error) {
	dialect, args, err := parseConfig(dep)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dialect, args)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MockGORM(_ module.Module) (*gorm.DB, error) {

	return gorm.Open("sqlite3", ":memory:")
}
