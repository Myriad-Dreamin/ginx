package probe

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type __PROBE___ struct {}
var __PROBE__ __PROBE___
var RootPackage string
var RootPath string

func init() {
	var InstancePackage, InstancePath string
	InstancePackage = reflect.TypeOf(__PROBE__).PkgPath()
	RootPackage = strings.ReplaceAll(filepath.Dir(InstancePackage), "\\", "/")
	_, InstancePath, _, _ = runtime.Caller(0)
	InstancePath = filepath.Dir(InstancePath)
	RootPath = filepath.Dir(InstancePath)
}
