package dblayer

import (
	"database/sql"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var rawDB *sql.DB
var dormDB *dorm.DB

func GetInstance() *gorm.DB {
	return db
}

func GetRawInstance() *sql.DB {
	return rawDB
}

func GetDormInstance() *dorm.DB {
	return dormDB
}
