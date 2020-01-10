package dblayer

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
)

var db = new(gorm.DB)
var rawDB = new(*sql.DB)
var dormDB = new(dorm.DB)

func GetInstance() *gorm.DB {
	return db
}

func GetRawInstance() *sql.DB {
	return *rawDB
}

func GetDormInstance() *dorm.DB {
	return dormDB
}
