package crud_dao

import (
	"github.com/jinzhu/gorm"
)

func Delete(db *gorm.DB) func(d interface{}) (int64, error) {
	return func(d interface{}) (int64, error) {
		rdb := db.Delete(d)
		return rdb.RowsAffected, rdb.Error
	}
}
