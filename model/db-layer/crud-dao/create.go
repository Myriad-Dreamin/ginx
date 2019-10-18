package crud_dao

import (
	"github.com/jinzhu/gorm"
)

func Create(db *gorm.DB) func(d interface{}) (int64, error) {
	return func(d interface{}) (int64, error) {
		rdb := db.Create(d)
		return rdb.RowsAffected, rdb.Error
	}
}
