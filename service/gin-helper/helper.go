package ginhelper

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/lib/errorc"
	"github.com/Myriad-Dreamin/minimum-template/lib/jwt"
	"github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/tidwall/gjson"
	"net/http"
	"reflect"
	"strconv"
)

var ResponseOK = serial.Response{Code: types.CodeOK}

func CheckInsertError(c controller.MContext, err error) bool {
	if code, errs := errorc.CheckInsertError(err); code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{Code: code, Error: errs})
		return true
	}
	return false
}

func MissID(c controller.MContext) {
	c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
		Code:  types.CodeInvalidParameters,
		Error: "id missing in the path",
	})
}

func AuthFailed(c controller.MContext, errorString string) {
	c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
		Code:  types.CodeAuthenticatePolicyError,
		Error: errorString,
	})
}

func ParseUint(c controller.MContext, key string) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative id",
		})
		return 0, false
	}
	return uint(id), true
}

func BindRequest(c controller.MContext, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return false
	}
	return true
}

func RawJson(c controller.MContext) (gjson.Result, bool) {
	if b, err := c.GetRawData(); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return gjson.Result{}, false
	} else {
		return gjson.ParseBytes(b), true
	}
}

func ParseUintAndBind(c controller.MContext, key string, req interface{}) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative id",
		})
		return 0, false
	}
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: err.Error(),
		})
		return 0, false
	}
	return uint(id), true
}

func RosolvePageVariable(c controller.MContext) (int, int, bool) {
	spage, ok := c.GetQuery("page")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "missing page number",
		})
		return 0, 0, false
	}
	page, err := strconv.Atoi(spage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeUnserializeDataError,
			Error: "can not convert page number to integer",
		})
		return 0, 0, false
	}
	spageSize, ok := c.GetQuery("page_size")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "missing page size",
		})
		return 0, 0, false
	}
	pageSize, err := strconv.Atoi(spageSize)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeUnserializeDataError,
			Error: "can not convert page size to integer",
		})
		return 0, 0, false
	}
	if page <= 0 || pageSize <= 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative params",
		})
		return 0, 0, false
	}
	return page, pageSize, true
}

func MaybeGetRawDataError(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeGetRawDataError,
			Error: err.Error(),
		})
		return true
	}
	return false
}

func MaybeCountError(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}

	return false
}

type applyContext struct{ controller.MContext }

func (ctx applyContext) applyError(code errorc.Code, errs string) bool {
	if code != types.CodeOK {
		ctx.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
			Code:  code,
			Error: errs,
		})
		return true
	}
	return false
}

func MaybeSelectError(c controller.MContext, anyObj interface{}, err error) bool {
	return applyContext{c}.applyError(errorc.MaybeSelectError(anyObj, err))
}

func MaybeSelectErrorWithTip(c controller.MContext, anyObj interface{}, err error, missError string) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if reflect.ValueOf(anyObj).IsNil() {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeNotFound,
			Error: missError,
		})
		return true
	}

	return false
}

func MaybeMissingError(c controller.MContext, has bool, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &serial.Response{
			Code: types.CodeNotFound,
		})
		return true
	}

	return false
}

func MaybeMissingErrorWithTip(c controller.MContext, has bool, err error, missError string) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeNotFound,
			Error: missError,
		})
		return true
	}

	return false
}
func MaybeOnlySelectError(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}

	return false
}

type Deletable interface {
	Delete() (int64, error)
}

func DeleteObj(c controller.MContext, deleteObj Deletable) bool {
	affected, err := deleteObj.Delete()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.Response{
			Code: types.CodeDeleteNoEffect,
		})
		return false
	}
	return true
}

func CreateObj(c controller.MContext, createObj errorc.Creatable) bool {
	if code, errs := errorc.CreateObj(createObj); code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{Code: code, Error: errs})
		return false
	}
	return true
}

func CreateObjWithTip(c controller.MContext, createObj errorc.Creatable) bool {
	if code, errs := errorc.CreateObj(createObj); code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{Code: code, Error: fmt.Sprintf("create %T failed: %v", createObj, errs)})
		return false
	}
	return true

}

type Updatable interface {
	Update() (int64, error)
}

func UpdateObj(c controller.MContext, updateObj Updatable) bool {
	affected, err := updateObj.Update()
	if err != nil {
		if CheckInsertError(c, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.Response{
			Code: types.CodeUpdateError,
		})
		return false
	}
	return true
}

func UpdateFields(c controller.MContext, obj errorc.UpdateFieldsable, fields []string) bool {
	return !applyContext{c}.applyError(errorc.UpdateFields(obj, fields))
}

func GetCustomFields(c controller.MContext) *types.CustomFields {
	claims, _ := c.Get("claims")
	return claims.(*jwt.CustomClaims).CustomField.(*types.CustomFields)
}
