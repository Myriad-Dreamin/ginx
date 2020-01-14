package artist

import (
	"reflect"
)

type norm struct {
	name  string
	tags  []*tag
	field Field
	param reflect.Value
}

type pureField struct {
	s string
}

func (p pureField) String() string {
	return p.s
}

func (n *norm) CreateParameterDescription(ctx *Context) ParameterDescription {
	desc := new(parameterDescription)
	desc.name = n.name
	desc.field = n.field
	if desc.field == nil {
		desc.field = pureField{fromSnakeToBigCamel(desc.name)}
	}

	if embedType, ok := n.param.Interface().(SerializeObject); ok {
		objDesc := embedType.CreateObjectDescription(ctx)
		desc.embedObjects = append(desc.embedObjects, objDesc)
		desc.pType = objDesc.GetType()
	} else {
		desc.pType = parseParamType(ctx, n)
		desc.source = parseSource(ctx, n)
		ctx.appendPackage(getReflectElementType(n.param).PkgPath())
	}
	desc.tags = make(map[string]string)
	desc.tags["json"] = desc.name
	desc.tags["form"] = desc.name
	for _, tag := range n.tags {
		if v, ok := desc.tags[tag.key]; ok {
			desc.tags[tag.key] = v + ";" + tag.value
		} else {
			desc.tags[tag.key] = tag.value
		}
	}

	return desc
}

type source struct {
	modelName  string
	faz        reflect.Type
	fazElem    reflect.Type
	fieldIndex int
}

func (s source) paramName() string {
	if len(s.modelName) != 0 {
		return s.modelName
	}
	return "_" + toSmallCamel(s.fazElem.Name())
}

func (s source) memberName() string {
	return s.fazElem.Field(s.fieldIndex).Name
}

func parseSource(context *Context, n *norm) *source {
	return context.getSource(n.param.UnsafeAddr())
}

type reflectType struct {
	reflect.Type
}

func parseParamType(ctx *Context, n *norm) Type {
	t := n.param.Type()
	if t != nil {
		ctx.appendPackage(t.PkgPath())
		return reflectType{t}
	} else {
		panic("nil type")
	}
}

type tag struct {
	key   string
	value string
}
type FieldName string

func createNorm(name string, descriptions ...interface{}) *norm {
	param := newNorm(name)
	for _, description := range descriptions {
		switch desc := description.(type) {
		case *tag:
			param.tags = append(param.tags, desc)
		case FieldName:
			param.field = pureField{string(desc)}
		case SerializeObject:
			param.param = reflect.ValueOf(desc)
		default:
			param.param = reflect.ValueOf(desc).Elem()
		}
	}
	return param
}

func Param(name string, descriptions ...interface{}) Parameter {
	return createNorm(name, descriptions...)
}

func Tag(key, value string) *tag {
	return &tag{
		key:   key,
		value: value,
	}
}

func newNorm(name string) *norm {
	return &norm{name: name}
}
