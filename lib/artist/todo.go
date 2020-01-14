package artist

import "fmt"

//func (desc objectDescription) String() string {
//	return fmt.Sprintf(`
//type %s struct {
//%s
//}
//
//%s
//
//%s
//
//%s
//`, desc.name, genStructFields(desc.params), desc.genPMethod(), desc.genMethod(), desc.genPackMethod())
//}

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
		res += "_" + fromSnakeToSmallCamel(param.GetDTOName()) + " " + param.GetType().String()
	}
	return res
}

type XParam struct {
	name, typeOf string
	source       ParameterDescription
}

func fromXParam(p []*XParam) (res string) {
	var appended = make(map[string]bool)
	for _, param := range p {
		if param.source.GetSource() != nil {
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

func fromFXParam(p []*XParam) (res string) {
	var appended = make(map[string]bool)
	for _, param := range p {
		if param.source.GetSource() != nil {
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

func fromSXParam(p []*XParam) (res string) {
	var appended = make(map[string]bool)
	for _, param := range p {
		if param.source.GetSource() != nil {
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

func (desc objectDescription) genResultFields(x []*XParam) (res string) {
	for i := range desc.params {
		param := desc.params[i]
		if len(res) != 0 {
			res += "\n"
		}
		res += "        " + param.GetField().String() + ": " + search(x, param) + ","
	}
	return
}
