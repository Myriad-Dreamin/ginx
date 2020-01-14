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

func (i transferClass) GenObjectTemplate() ObjectTemplate {
	x := &XParam{
		name:   fromBigCamelToSnake(getReflectTypeElementType(i.baseType).Name()),
		typeOf: i.baseType.String(),
		source: nil,
	}
	f := &ObjectTemplateFieldImpl{
		Name:   "",
		PType:  i.baseType,
		Tag:    nil,
		Source: x,
	}
	return &ObjectTemplateImpl{
		Name:   i.name,
		TType:  TemplateTypeEq,
		Fields: []ObjectTemplateField{f},
		Xps:    []*XParam{x},
	}
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
