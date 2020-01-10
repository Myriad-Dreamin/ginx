package dblayer

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
	"github.com/Myriad-Dreamin/minimum-template/lib/core"
	"github.com/Myriad-Dreamin/minimum-template/lib/fcg"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
)


func GetInstance() *gorm.DB {
	return p.GetInstance()
}

func GetRawInstance() *sql.DB {
	return p.GetRawInstance()
}

func GetDormInstance() *dorm.DB {
	return p.GetDormInstance()
}

func Register(m module.Module)  error {
	return p.Register(m)
}

func Configuration(cfg *config.ServerConfig) {
	(*p.RawDB).SetMaxIdleConns(cfg.DatabaseConfig.MaxIdle)
	(*p.RawDB).SetMaxOpenConns(cfg.DatabaseConfig.MaxActive)
}

type Traits = traits.ModelTraits
type Interface = traits.Interface
type TraitsAcceptObject = traits.ORMObject
type where1Func = func(interface{}) (interface{}, error)

func NewTraits(t TraitsAcceptObject) Traits {
	tt := traits.NewModelTraits(t, p.DB, p.DormDB)
	return tt
}

var p = mcore.NewModelModule(_module{})
type _module struct{}

func (_module) Migrates() error {
	return fcg.Calls([]fcg.MaybeInitor{
		User{}.migrate,
	})
}

func (_module) Injects() error {
	return fcg.Calls([]fcg.MaybeInitor{
		injectUserTraits,
	})
}




