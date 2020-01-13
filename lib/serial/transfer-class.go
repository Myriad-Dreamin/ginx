package serial

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
	name string
	base interface{}
}

func (i transferClass) String() string {
	return fmt.Sprintf("type %s = %s", i.name, reflect.TypeOf(i.base))
}

func (i transferClass) GetEmbedObject() []ObjectDescription {
	return nil
}

func Transfer(name string, base interface{}) *transferClass {
	return &transferClass{name:name, base: base}
}

