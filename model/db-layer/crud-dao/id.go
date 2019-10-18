package crud_dao

import "github.com/jinzhu/gorm"

func ID(db *gorm.DB) func(obj interface{}, id uint) (err error) {
	return func(obj interface{}, id uint) (err error) {
		rdb := db.First(obj, id)
		err = rdb.Error
		if rdb.RecordNotFound() {
			obj = nil
			err = nil
		}
		return
	}
}

