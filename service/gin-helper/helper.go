package ginhelper

import (
	"fmt"
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"reflect"
	"strconv"
)

type Response struct {
	Code int `json:"code"`
}

type ErrorSerializer struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

var ResponseOK = Response{Code: types.CodeOK}

func CheckInsertError(c *gin.Context, err error) bool {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			c.AbortWithStatusJSON(http.StatusOK, &Response{Code: types.CodeDuplicatePrimaryKey})
			return true
		}
	}
	return false
}

func MissID(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
		Code:  types.CodeBindError,
		Error: "id missing in the path",
	})
}

func AuthFailed(c *gin.Context, errorString string) {
	c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
		Code:  types.CodeAuthenticatePolicyError,
		Error: errorString,
	})
}

func ParseUint(c *gin.Context, key string) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeBindError,
			Error: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative id",
		})
		return 0, false
	}
	return uint(id), true
}

func BindRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeBindError,
			Error: err.Error(),
		})
		return false
	}
	return true
}

func ParseUintAndBind(c *gin.Context, key string, req interface{}) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeBindError,
			Error: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative id",
		})
		return 0, false
	}
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeBindError,
			Error: err.Error(),
		})
		return 0, false
	}
	return uint(id), true
}

func RosolvePageVariable(c *gin.Context) (int, int, bool) {
	spage, ok := c.GetQuery("page")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeBindError,
			Error: "missing page number",
		})
		return 0, 0, false
	}
	page, err := strconv.Atoi(spage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeUnserializeDataError,
			Error: "can not convert page number to integer",
		})
		return 0, 0, false
	}
	spageSize, ok := c.GetQuery("page_size")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeBindError,
			Error: "missing page size",
		})
		return 0, 0, false
	}
	pageSize, err := strconv.Atoi(spageSize)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeUnserializeDataError,
			Error: "can not convert page size to integer",
		})
		return 0, 0, false
	}
	if page <= 0 || pageSize <= 0 {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeInvalidParameters,
			Error: "bad negative params",
		})
		return 0, 0, false
	}
	return page, pageSize, true
}

func MaybeGetRawDataError(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeGetRawDataError,
			Error: err.Error(),
		})
		return true
	}
	return false
}

func MaybeCountError(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}

	return false
}

func MaybeSelectError(c *gin.Context, anyObj interface{}, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if reflect.ValueOf(anyObj).IsNil() {
		c.AbortWithStatusJSON(http.StatusOK, &Response{
			Code: types.CodeNotFound,
		})
		return true
	}

	return false
}

func MaybeSelectErrorWithTip(c *gin.Context, anyObj interface{}, err error, missError string) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if reflect.ValueOf(anyObj).IsNil() {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeNotFound,
			Error: missError,
		})
		return true
	}

	return false
}

func MaybeMissingError(c *gin.Context, has bool, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &Response{
			Code: types.CodeNotFound,
		})
		return true
	}

	return false
}

func MaybeMissingErrorWithTip(c *gin.Context, has bool, err error, missError string) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeSelectError,
			Error: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeNotFound,
			Error: missError,
		})
		return true
	}

	return false
}
func MaybeOnlySelectError(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
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

func DeleteObj(c *gin.Context, deleteObj Deletable) bool {
	affected, err := deleteObj.Delete()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &ErrorSerializer{
			Code:  types.CodeDeleteError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &Response{
			Code: types.CodeDeleteNoEffect,
		})
		return false
	}
	return true
}

type Creatable interface {
	Create() (int64, error)
}

func CreateObj(c *gin.Context, createObj Creatable) bool {
	affected, err := createObj.Create()
	if err != nil {
		if CheckInsertError(c, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &Response{
			Code: types.CodeInsertError,
		})
		return false
	}
	return true
}

func CreateObjWithTip(c *gin.Context, createObj Creatable) bool {
	affected, err := createObj.Create()
	if err != nil {
		if CheckInsertError(c, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeInsertError,
			Error: fmt.Sprintf("create %T failed: %v", createObj, err.Error()),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &Response{
			Code: types.CodeInsertError,
		})
		return false
	}
	return true
}

type Updatable interface {
	Update() (int64, error)
}

func UpdateObj(c *gin.Context, updateObj Updatable) bool {
	affected, err := updateObj.Update()
	if err != nil {
		if CheckInsertError(c, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &Response{
			Code: types.CodeUpdateError,
		})
		return false
	}
	return true
}

type UpdateFieldsable interface {
	UpdateFields(fields []string) (int64, error)
}

func UpdateFields(c *gin.Context, obj UpdateFieldsable, fields []string) bool {
	_, err := obj.UpdateFields(fields)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorSerializer{
			Code:  types.CodeUpdateError,
			Error: err.Error(),
		})
		return false
	}
	return true
}

func GetCustomFields(c *gin.Context) *types.CustomFields {
	claims, _ := c.Get("claims")
	return claims.(*jwt.CustomClaims).CustomField.(*types.CustomFields)
}
