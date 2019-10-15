package router

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Checker func(c *gin.Context, enforceVariables ...interface{}) bool
type AuthFailedFunc func(c *gin.Context, errorString string)
type ValidateFunc func(c *gin.Context, sbj string) bool
type WrappedValidateFunc func(c *gin.Context) bool

type common struct {
	object     string
	placeholder string
	MissID      func(c *gin.Context)
	Check       Checker
}

type ObjectMiddleware struct {
	*common
	validate func(c *gin.Context, sbj, obj string) bool
}

func defaultMissID(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}
func defaultAuthFailed(c *gin.Context, errorString string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, errorString)
}

func NewObjectMiddleware(object, placeholder string, MissID gin.HandlerFunc, checker Checker) *ObjectMiddleware {
	if MissID == nil {
		MissID = defaultMissID
	}
	if checker == nil {
		panic("checker must not be null")
	}
	return &ObjectMiddleware{
		common: &common{
			object:     object + ":",
			placeholder: placeholder,
			MissID:      MissID,
			Check:       checker,
		},
		validate: func(c *gin.Context, sbj, obj string) bool { return true },
	}
}

func (objFac *ObjectMiddleware) ToObject(id uint) string {
	return objFac.object + strconv.Itoa(int(id))
}

func (objFac *ObjectMiddleware) ToObjectString(id string) string {
	return objFac.object + id
}

func (objFac *ObjectMiddleware) Validate() ValidateFunc {
	return func(c *gin.Context, sbj string) bool {
		if id := c.Param(objFac.placeholder); len(id) == 0 {
			objFac.MissID(c)
			return false
		} else if !objFac.validate(c ,sbj, objFac.ToObjectString(id)) {
			return false
		}
		return true
	}
}

func (objFac *ObjectMiddleware) Absorb(action string) {
	oldValidation := objFac.validate
	objFac.validate = func(c *gin.Context, subject, object string) bool {
		if !oldValidation(c, subject, object) {
			return false
		}
		//middleware.uTable+uid, c.Request.URL.Path, c.Request.Method
		if !objFac.Check(c, subject, object, action) {
			return false
		}
		return true
	}
}

type Validator interface {
	Enforce(...interface{}) (bool, error)
}

type commonX struct {
	uTable           string
	idKey            string
	reportMissID     gin.HandlerFunc
	authFailed       AuthFailedFunc
	checker          Checker
}

type Middleware struct {
	*commonX
	reqs Requirements
	objectMiddleware map[string]*ObjectMiddleware
}

func fromValidatorAndAuthFailed(validator Validator, AuthFailed AuthFailedFunc) Checker {
	return func(c *gin.Context, enforceVariables ...interface{}) bool {
		if ok, err := validator.Enforce(enforceVariables...); err != nil {
			AuthFailed(c, err.Error())
			return false
		} else if !ok {
			AuthFailed(c, fmt.Sprintf("no authentication %v", enforceVariables))
			return false
		}
		return true
	}
}

func newMiddlewareFromCommonX(reqs Requirements, commonX *commonX)  *Middleware {
	return &Middleware{
		objectMiddleware: make(map[string]*ObjectMiddleware),
		reqs: reqs,
		commonX: commonX,
	}
}

func NewMiddleware(validator Validator, uTable, idKey string, MissID gin.HandlerFunc, AuthFailed AuthFailedFunc) *Middleware {
	if AuthFailed == nil {
		AuthFailed = defaultAuthFailed
	}
	return &Middleware{
		objectMiddleware: make(map[string]*ObjectMiddleware),
		commonX: &commonX{
			uTable:           uTable,
			idKey:            idKey,
			reportMissID:     MissID,
			authFailed:       AuthFailed,
			checker:          fromValidatorAndAuthFailed(validator, AuthFailed),
		},
	}
}

func (mw *Middleware) Copy(reqs... Requirement) *Middleware {
	nmw := newMiddlewareFromCommonX(mw.reqs.PlusX(reqs), mw.commonX)
	for sbj, omw := range mw.objectMiddleware {
		nmw.objectMiddleware[sbj] = &ObjectMiddleware{
			common: omw.common,
			validate:    omw.validate,
		}
	}
	return nmw
}

func (mw *Middleware) Group(object, placeholder string, reqs... Requirement) *Middleware {
	nmw := mw.Copy(reqs...)
	nmw.objectMiddleware[object] = NewObjectMiddleware(object, placeholder, nmw.reportMissID, nmw.checker)
	return nmw
}

func (mw *Middleware) MaybeGroup(object, placeholder string, reqs... Requirement) (*Middleware, bool) {
	if _, ok := mw.objectMiddleware[object]; ok {
		return nil, false
	}
	return mw.Group(object, placeholder, reqs...), true
}

func (mw *Middleware) MustGroup(object, placeholder string, reqs... Requirement) *Middleware {
	if _, ok := mw.objectMiddleware[object]; !ok {
		return mw.Group(object, placeholder, reqs...)
	}
	panic(fmt.Errorf("cannot register %v(%v) object", object, placeholder))
}


func (mw *Middleware) useValidate(validate ValidateFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		subjectID := c.GetString(mw.idKey)
		if len(subjectID) == 0 {
			_ = c.AbortWithError(http.StatusForbidden, errors.New("missing uid"))
			return
		}
		if !validate(c, mw.uTable + subjectID) {
			return
		}
	}
}

type Requirement struct {
	Object string
	Action  string
}

type RetrieveFunc func(c *gin.Context) string

func (r Requirement) RetrieveFunc(placeholder string) RetrieveFunc {
	return func(c *gin.Context) string {
		if id := c.Param(placeholder); len(id) == 0 {
			return ""
		}
		return r.Object + ":" + placeholder
	}
}

type Requirements []Requirement

func (r Requirements) Plus(req Requirement) Requirements {
	return append(r, req)
}
func (r Requirements) PlusX(reqs Requirements) Requirements {
	return append(r, reqs...)
}

func (mw *Middleware) AdminOnlyValidate() ValidateFunc {
	return func(c *gin.Context, sbj string) bool {
		return mw.checker(c, sbj, "_", "auth")
	}
}

func (mw *Middleware) BuildValidate(reqs... Requirement) ValidateFunc {
	for _, req := range mw.reqs.PlusX(reqs) {
		if mw, ok := mw.objectMiddleware[req.Object]; !ok {
			panic(fmt.Errorf("missing requirement object, want %s, but not registered in the middleware", req.Object))
		} else {
			mw.Absorb(req.Action)
		}
	}

	validate := func(c *gin.Context, sbj string) bool {
		return true
	}

	for _, omw := range mw.objectMiddleware {
		nextValidate := omw.Validate()
		thisValidate := validate
		validate = func(c *gin.Context, sbj string) bool {
			if !thisValidate(c, sbj) {
				return false
			}
			return nextValidate(c, sbj)
		}
	}

	return validate
}

func (mw *Middleware) Wrap(validate ValidateFunc) WrappedValidateFunc {
	return func(c *gin.Context) bool {
		subjectID := c.GetString(mw.idKey)
		if len(subjectID) == 0 {
			_ = c.AbortWithError(http.StatusForbidden, errors.New("missing uid"))
			return false
		}
		if !validate(c, mw.uTable + subjectID) {
			return false
		}
		return true
	}
}

func (mw *Middleware) AdminOnly() gin.HandlerFunc {
	return mw.useValidate(mw.AdminOnlyValidate())
}

func (mw *Middleware) Build(reqs... Requirement) gin.HandlerFunc {
	return mw.useValidate(mw.BuildValidate(reqs...))
}


