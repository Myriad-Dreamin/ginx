package crud_dao

import (
	"github.com/jinzhu/gorm"
)

func Update(db *gorm.DB) func(d interface{}) (int64, error) {
	return func(d interface{}) (int64, error) {
		rdb := db.Save(d)
		return rdb.RowsAffected, rdb.Error
	}
}
