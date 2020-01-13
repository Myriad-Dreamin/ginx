package serial

import (
	"fmt"
	"unicode"
)

type SerializeObjectI interface {
	CreateObjectDescription(ctx *Context) ObjectDescription
}


type objectDescription struct {
	name string
	params []*ParameterDescription
}

func (desc objectDescription) GetTypeString() string {
	return desc.name
}

func (desc objectDescription) GetEmbedObject() (dx []ObjectDescription) {
	for _, param := range desc.params {
		dx = append(dx, param.embedObjects...)
	}
	return dx
}

func (desc objectDescription) String() string {
	return fmt.Sprintf(`
type %s struct {
%s
}

%s

%s

%s
`, desc.name, genStructFields(desc.params),desc.genPMethod(), desc.genMethod(), desc.genPackMethod() )
}

func (desc objectDescription) genPMethod() string {
	xs := desc.genXParams()
	return fmt.Sprintf(`func PSerialize%s(%s) *%s {
	return &%s{
%s
	}
}`, desc.name,
		fromXParam(xs), desc.name, desc.name, desc.genResultFields(xs))
}

func (desc objectDescription) genMethod() string {
	xs := desc.genXParams()
	return fmt.Sprintf(`func Serialize%s(%s) %s {
	return %s{
%s
	}
}`, desc.name,
		fromXParam(xs), desc.name, desc.name, desc.genResultFields(xs))
}

func (desc objectDescription) genPackMethod() string {
	xs := desc.genXParams()
	if len(xs) == 0 {
		return fmt.Sprintf(`func PackSerialize%s() (pack []%s) {
	return
}`, desc.name, desc.name)
	}
	return fmt.Sprintf(`func PackSerialize%s(%s) (pack []%s) {
	for i := range %s {
		pack = append(pack, Serialize%s(%s))
	}
	return
}`, desc.name, fromSXParam(xs), desc.name, xs[0].name, desc.name, fromFXParam(xs))
}

func (desc objectDescription) genParams() (res string) {
	//desc.params
	for _, param := range desc.params {
		if len(res) != 0 {
			res += ", "
		}
		res += "_" + fromSnakeToSmallCamel(param.name) + " " + param.typeString
	}
	return res
}

type xParam struct {
	name, typeOf string
	source *ParameterDescription
}

func (desc objectDescription) genXParams() (params []xParam) {
	//desc.params
	for _, param := range desc.params {
		if param.source != nil {
			params = append(params, xParam{
				name:   param.source.paramName(),
				typeOf: param.source.faz.String(),
				source: param,
			})
		} else {
			params = append(params, xParam{
				name:   "_" + fromSnakeToSmallCamel(param.name),
				typeOf: param.typeString,
				source: param,
			})
		}
	}
	return
}

func toSmallCamel(name string) string {
	if len(name) == 0 {
		return name
	} else {
		return string(unicode.ToLower(rune(name[0]))) + name[1:]
	}
}

func fromXParam(p []xParam) (res string) {
	var appended = make(map[string]bool)
	for _, param := range p {
		if param.source.source != nil {
			if _, ok := appended[param.name]; ok {
				continue
			}
			appended[param.name] = true
		}
		if len(res) != 0 {
			res += ", "
		}
		res += param.name + " " + param.typeOf
	}
	return res
}

func fromFXParam(p []xParam) (res string) {
	var appended = make(map[string]bool)
	for _, param := range p {
		if param.source.source != nil {
			if _, ok := appended[param.name]; ok {
				continue
			}
			appended[param.name] = true
		}
		if len(res) != 0 {
			res += ", "
		}
		res += param.name + "[i]"
	}
	return res
}

func fromSXParam(p []xParam) (res string) {
	var appended = make(map[string]bool)
	for _, param := range p {
		if param.source.source != nil {
			if _, ok := appended[param.name]; ok {
				continue
			}
			appended[param.name] = true
		}
		if len(res) != 0 {
			res += ", "
		}
		res += param.name + " []" + param.typeOf
	}
	return res
}

func (desc objectDescription) genResultFields(x []xParam) (res string) {
	for i := range desc.params {
		param := desc.params[i]
		if len(res) != 0 {
			res += "\n"
		}
		res += "        " + param.fieldName + ": " + search(x, param) + ","
	}
	return
}

func search(params []xParam, sp *ParameterDescription) string {
	for _, param := range params {
		if sp == param.source {
			if sp.source == nil {
				return param.name
			} else {
				return param.name + "." + sp.source.fazElem.Field(sp.source.fieldIndex).Name
			}
		}
	}
	panic("not found")
}




type ObjectDescription interface {
	fmt.Stringer
	GetTypeString() string
	GetEmbedObject() []ObjectDescription
}


func dumpObj(res string, desc ObjectDescription) (result string) {
	result = res
	result += "\n" + desc.String()
	for _, obj := range desc.GetEmbedObject() {
		result = dumpObj(result, obj)
	}
	return
}

