package errorc

import (
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/go-sql-driver/mysql"
	"reflect"
	"strconv"
)

type Code = types.CodeRawType

func MaybeSelectError(anyObj interface{}, err error) (Code, string) {
	if err != nil {
		return types.CodeSelectError, err.Error()
	}
	if reflect.ValueOf(anyObj).IsNil() {
		return types.CodeNotFound, "not found"
	}
	return types.CodeOK, ""
}

type UpdateFieldsable interface {
	UpdateFields(fields []string) (int64, error)
}

func UpdateFields(obj UpdateFieldsable, fields []string) (Code, string) {
	_, err := obj.UpdateFields(fields)
	if err != nil {
		return types.CodeUpdateError, err.Error()
	}
	return types.CodeOK, ""
}

type Creatable interface {
	Create() (int64, error)
}

func CheckInsertError(err error) (Code, string) {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		switch mysqlError.Number {
		case 1062:
			return types.CodeDuplicatePrimaryKey, ""
		case 1366:
			return types.CodeDatabaseIncorrectStringValue, ""
		default:
			return types.CodeInsertError, strconv.Itoa(int(mysqlError.Number))
		}
	}
	return types.CodeOK, ""
}

func CreateObj(createObj Creatable) (Code, string) {
	affected, err := createObj.Create()
	if err != nil {
		if code, errs := CheckInsertError(err); code != types.CodeOK {
			return code, errs
		}
	} else if affected == 0 {
		return types.CodeInsertError, "affect nothing"
	}
	return types.CodeOK, ""
}
