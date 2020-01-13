package serial

import (
	"reflect"
)

type norm struct {
	name      string
	tags      []*tag
	fieldName string
	param     reflect.Value
}

func (n *norm) Create(ctx *Context) *ParameterDescription {
	desc := new(ParameterDescription)
	desc.name = n.name
	desc.fieldName = n.fieldName
	if len(desc.fieldName) == 0 {
		desc.fieldName = fromSnakeToBigCamel(desc.name)
	}

	if embedType, ok := n.param.Interface().(SerializeObjectI); ok {
		objDesc := embedType.CreateObjectDescription(ctx)
		desc.embedObjects = append(desc.embedObjects, objDesc)
		desc.typeString = objDesc.GetTypeString()
	} else {
		desc.typeString = parseParamType(ctx, n)
		desc.source = parseSource(ctx, n)
		t := n.param.Type()
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		pk := t.PkgPath()
		if len(pk) != 0 {
			ctx.appendPackage(pk)
		}
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

func parseSource(context *Context, n *norm) *source {
	return context.getSource(n.param.UnsafeAddr())
}

func parseParamType(ctx *Context, n *norm) string {
	t := n.param.Type()
	if t != nil {
		ctx.appendPackage(t.PkgPath())
		return t.String()
	} else {
		panic("nil type")
	}
}

type tag struct {
	key   string
	value string
}
type FieldName string

func Param(name string, descriptions ...interface{}) Parameter {
	param := newNorm(name)
	for _, description := range descriptions {
		switch desc := description.(type) {
		case *tag:
			param.tags = append(param.tags, desc)
		case FieldName:
			param.fieldName = string(desc)
		case SerializeObjectI:
			param.param = reflect.ValueOf(desc)
		default:
			param.param = reflect.ValueOf(desc).Elem()
		}
	}
	return param
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
