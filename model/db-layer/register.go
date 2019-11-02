package dblayer

import (
	"github.com/Myriad-Dreamin/core-oj/config"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/ginx/types"
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

func Register(rdb *gorm.DB, logger types.Logger) error {
	var err error
	*db = *rdb.Debug()
	*rawDB = *db.DB()

	if err = rawDB.Ping(); err != nil {
		return err
	}

	xdb, err := dorm.FromRaw(rawDB, adapt(logger))
	if err != nil {
		return err
	}

	*dormDB = *xdb

	return migrates()
}

func Configuration(cfg *config.Configuration) {

	rawDB.SetMaxIdleConns(100)
	rawDB.SetMaxOpenConns(100)
}
