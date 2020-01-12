package serial

import (
	"fmt"
)

type ParameterDescription struct {
	embedObjects []*objectDescription
	TypeString string
	name       string
	fieldName  string
	tags       map[string]string
}

type objectDescription struct {
	name string
	params []*ParameterDescription
}

func (description objectDescription) String() string {
	return fmt.Sprintf(`
type %s struct {
%s
}`, description.name, genStructFields(description.params))
}

func genStructFields(descriptions []*ParameterDescription) (res string) {
	for i := range descriptions {
		desc := descriptions[i]
		if len(res) != 0 {
			res += "\n"
		}
		res += "    " + desc.fieldName + " " + desc.TypeString + genTag(desc.tags)
	}
	return
}

func genTag(tags map[string]string) (res string) {
	res = "`"
	for k, v := range tags {
		res += fmt.Sprintf(`%s:"%s"`, k, v)
	}
	res += "`"
	return res
}

type methodDescription struct {
	method MethodType
	name   string
	requests []*objectDescription
	replies []*objectDescription
}

type categoryDescription struct {
	path     string
	methods  []*methodDescription
	packages map[string]int
}

type serviceDescription struct {
	name       string
	categories []*categoryDescription
	filePath   string
	packages   map[string]int
}

func (description serviceDescription) generateObjects() (result string) {
	for _, cat := range description.categories {
		for _, method := range cat.methods {
			for _, req := range method.requests {
				result += "\n" + req.String()
				for _, param := range req.params {
					for _, obj := range param.embedObjects {
						result += "\n" + obj.String()
					}
				}
			}

			for _, res := range method.replies {
				result += "\n" + res.String()
				for _, param := range res.params {
					for _, obj := range param.embedObjects {
						result += "\n" + obj.String()
					}
				}
			}
		}
	}
	return
}

