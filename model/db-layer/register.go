package dblayer

import (
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
)

func migrates() error {
	for _, migrate := range []func() error{
		User{}.migrate,
	} {
		if err := migrate(); err != nil {
			return err
		}
	}
	return nil
}

func Register(rdb *gorm.DB, m module.Module) error {
	var err error
	*db = *rdb
	*rawDB = db.DB()

	// test if it is alive
	if err = (*rawDB).Ping(); err != nil {
		return err
	}

	xdb, err := dorm.FromRaw(*rawDB,
		adapt(m.Require(config.ModulePath.Global.Logger).(types.Logger)),
		dorm.Escaper(
			m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig).DatabaseConfig.Escaper))
	if err != nil {
		return err
	}

	*dormDB = *xdb

	return migrates()
}

func Configuration(cfg *config.ServerConfig) {

	(*rawDB).SetMaxIdleConns(cfg.DatabaseConfig.MaxIdle)
	(*rawDB).SetMaxOpenConns(cfg.DatabaseConfig.MaxActive)
}
