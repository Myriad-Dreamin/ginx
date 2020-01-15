package artisan

import "reflect"

// todo export
type source struct {
	modelName  string
	faz        reflect.Type
	fazElem    reflect.Type
	fieldIndex int
}

func (s source) ParamName() string {
	if len(s.modelName) != 0 {
		return s.modelName
	}
	return "_" + toSmallCamel(s.fazElem.Name())
}

func (s source) MemberName() string {
	return s.fazElem.Field(s.fieldIndex).Name
}
