package artist

import (
	"fmt"
	"reflect"
)

func (i transferClass) CreateObjectDescription(ctx *Context) ObjectDescription {
	ctx.appendPackage(reflect.TypeOf(i.base).PkgPath())
	return i
}

func (i transferClass) GetTypeString() string {
	return i.name
}

type transferClass struct {
	name     string
	base     interface{}
	baseType reflect.Type
}

func (i transferClass) GetType() Type {
	return i.baseType
}

func (i transferClass) String() string {
	return fmt.Sprintf("type %s = %s", i.name, i.baseType)
}

func (i transferClass) GetEmbedObject() []ObjectDescription {
	return nil
}

func Transfer(name string, base interface{}) *transferClass {
	return &transferClass{name: name, base: base, baseType: reflect.TypeOf(base)}
}
