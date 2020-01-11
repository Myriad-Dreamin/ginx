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

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
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

func FromContext(dep module.Module) bool {
	return p.FromContext(dep)
}

func Install(dep module.Module) bool {
	return p.Install(dep)
}

func InstallMock(dep module.Module) bool {
	return p.InstallMock(dep)
}

func Close(dep module.Module) bool {
	if err := p.GormDB.Close(); err != nil {
		p.Logger.Error("close DB error", "error", err)
		return false
	}
	return true
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

func (m *modelModule) install(
	initFunc func(dep module.Module) bool, dep module.Module) bool {
	return true && 
		m.LoggerModule.Install(dep) &&
		initFunc(dep) &&
		m.RawSQLModule.FromRaw(m.GormDB.DB(), dep) &&
		m.DormModule.FromRawSQL(m.RawDB, dep) && mcore.ModelCallback(m, dep)
}

func (m *modelModule) FromContext(dep module.Module) bool {
	return m.install(m.GormModule.FromContext, dep)
}

func (m *modelModule) Install(dep module.Module) bool {
	return m.install(m.GormModule.InstallFromConfiguration, dep)
}

func (m *modelModule) InstallMock(dep module.Module) bool {
	return m.install(m.GormModule.InstallMockFromConfiguration, dep)
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
