package dblayer

import (
	"github.com/Myriad-Dreamin/core-oj/config"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func migrates() error {
	for _, migrate := range []func() error {
		// do not delete following line
		//Object{}.migrate,
		Submission{}.migrate,
	} {
		if err := migrate(); err != nil {
			return err
		}
	}
	return nil
}

func Register(rdb *gorm.DB, logger types.Logger) error {
	var err error
	if db == nil {
		db = new(gorm.DB)
	}
	*db = *rdb
	rawDB = db.DB()

	if err = rawDB.Ping(); err != nil {
		return err
	}

	dormDB, err = dorm.FromRaw(rawDB, adapt(logger))
	if err != nil {
		return err
	}

	return migrates()
}

func Configuration(cfg *config.Configuration) {

	rawDB.SetMaxIdleConns(100)
	rawDB.SetMaxOpenConns(100)
}


