package dblayer

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/lib/core"
	"github.com/Myriad-Dreamin/minimum-template/lib/fcg"
	"github.com/jinzhu/gorm"
)

var p = newModelModule()

func GetInstance() *gorm.DB {
	return p.GetGormInstance()
}

func GetRawInstance() *sql.DB {
	return p.GetRawSQLInstance()
}

func GetDormInstance() *dorm.DB {
	return p.GetDormInstance()
}

func Register(m module.Module) bool {
	return p.Install(m)
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
	tt := traits.NewModelTraits(t, p.GormDB, p.DormDB)
	return tt
}

type modelModule struct {
	mcore.GormModule
	mcore.RawSQLModule
	mcore.DormModule
	mcore.LoggerModule
}

func newModelModule() modelModule {
	return modelModule{}
}

func (m *modelModule) Install(dep module.Module) bool {
	return true && 
		m.LoggerModule.Install(dep) &&
		m.GormModule.FromContext(dep) &&
		m.RawSQLModule.FromRaw(m.GormDB.DB(), dep) &&
		m.DormModule.FromRawSQL(m.RawDB, dep) && mcore.ModelCallback(m, dep)
}
// var p = &P

func (modelModule) Migrates() error {
	return fcg.Calls([]fcg.MaybeInitor{
		User{}.migrate,
	})
}

func (modelModule) Injects() error {
	return fcg.Calls([]fcg.MaybeInitor{
		injectUserTraits,
	})
}
