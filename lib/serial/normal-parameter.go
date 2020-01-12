package serial

import (
	"bytes"
	"reflect"
	"unicode"
)

type norm struct {
	name string
	tags []*tag
	fieldName string
	paramType interface{}
}

func fromSnake(src string) string {
	if len(src) == 0 {
		return ""
	}
	var b = bytes.NewBuffer(make([]byte, 0, len(src)))
	var bo = true
	for i := range src {
		if src[i] == '_' {
			bo = true
		} else {
			if bo {
				bo = false
				b.WriteByte(byte(unicode.ToUpper(rune(src[i]))))
			} else {
				b.WriteByte(src[i])
			}
		}
	}
	return b.String()
}

func (n *norm) Create(ctx *context) *ParameterDescription {
	desc := new(ParameterDescription)
	desc.name = n.name
	desc.fieldName = n.fieldName
	if len(desc.fieldName) == 0 {
		desc.fieldName = fromSnake(desc.name)
	}

	if embedType, ok := n.paramType.(interface{
		create(ctx *context) *objectDescription
	}); ok {
		objDesc := embedType.create(ctx)
		desc.embedObjects = append(desc.embedObjects, objDesc)
		desc.TypeString = objDesc.name
	} else {
		desc.TypeString = parseParamType(ctx, n)
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

func parseParamType(ctx *context, n *norm) string {
	//fmt.Println(svc, n)
	switch v := n.paramType.(type) {
	case *SerializeObject:
	default:
		t := reflect.TypeOf(v)
		if t != nil {
			ctx.appendPackage(t.PkgPath())
			return t.String()
		} else {
			panic("nil type")
		}
	}
	return "todo"
}

type tag struct {
	key string
	value string
}
type FieldName string


func Param(name string, descriptions...interface{}) Parameter {
	param := newNorm(name)
	for _, description := range descriptions {
		switch desc := description.(type) {
		case *tag:
			param.tags = append(param.tags, desc)
		case FieldName:
			param.fieldName = string(desc)
		default:
			param.paramType = desc
		}
	}
	return param
}

func Tag(key, value string) *tag {
	return &tag{
		key: key,
		value: value,
	}
}

func newNorm(name string) *norm {
	return &norm{name:name}
}

