package serial

import "reflect"

type inheritClass struct {
	name string
	bases []interface{}
}

func (i inheritClass) Create(svc ProposingService) *ParameterDescription {
	desc := new(ParameterDescription)
	if len(i.name) == 0 {
		base := i.bases[0]
		t := reflect.TypeOf(base)
		desc.name = ""
		desc.TypeString = t.Name()
	} else {
		//desc.embedObjects = append()
		panic("todo")
	}
	return desc
}

func Transfer(base interface{}) *inheritClass {
	return &inheritClass{name:"", bases:[]interface{}{base}}
}

func Inherit(name string, bases ...interface{}) *inheritClass {
	return &inheritClass{name:name, bases:bases}
}

