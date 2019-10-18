package crud_dao

import (
	"github.com/Myriad-Dreamin/dorm"
)




func UpdateFields(objectModel *dorm.Model) func(d dorm.ORMObject, fields []string) (int64, error) {
	return func(d dorm.ORMObject, fields []string) (int64, error) {
		return objectModel.Anchor(d).Select(fields...).UpdateFields()
	}
}

