package artist

import "fmt"

type ObjectTemplateMiddleware = func(template ObjectTemplate, addition TemplateContext) ObjectTemplate

// TemplateType = {Struct, Interface}
// type :TemplateName :TemplateType {
//     ...:TemplateFields
//            :Name :Type :Tag :SourceXParams
// }
//
// func method(:XParams)
//      XParam:
//          :Name :Type

// type TemplateName TemplateType(=) TemplateFields[0].Type

type TemplateType int

const (
	TemplateTypeStruct TemplateType = iota
	TemplateTypeInterface
	TemplateTypeEq
)

type ObjectTemplateField interface {
	GetName() string
	GetType() Type
	GetTag() map[string]string
	GetSource() *XParam
}

type ObjectTemplate interface {
	fmt.Stringer
	GetName() string
	GetType() TemplateType
	GetFields() []ObjectTemplateField
	GetSources() []*XParam
}

type ObjectTemplateImpl struct {
	Name   string
	TType  TemplateType
	Fields []ObjectTemplateField
	Xps    []*XParam
}

func (o ObjectTemplateImpl) String() string {
	switch o.TType {
	case TemplateTypeEq:
		return fmt.Sprintf(`
type %s = %s`, o.Name, o.Fields[0].GetType().String())
	case TemplateTypeStruct:
		return fmt.Sprintf(`
type %s struct {
%s
}`, o.Name, instantiateStructFields(o.Fields))
	default:
		panic("todo")
	}
}

func (o ObjectTemplateImpl) GetName() string {
	return o.Name
}

func (o ObjectTemplateImpl) GetSources() []*XParam {
	return o.Xps
}

func (o ObjectTemplateImpl) GetType() TemplateType {
	return o.TType
}

func (o ObjectTemplateImpl) GetFields() []ObjectTemplateField {
	return o.Fields
}

type ObjectTemplateFieldImpl struct {
	Name   string
	PType  Type
	Tag    map[string]string
	Source *XParam
}

func (o ObjectTemplateFieldImpl) GetName() string {
	return o.Name
}

func (o ObjectTemplateFieldImpl) GetType() Type {
	return o.PType
}

func (o ObjectTemplateFieldImpl) GetTag() map[string]string {
	return o.Tag
}

func (o ObjectTemplateFieldImpl) GetSource() *XParam {
	return o.Source
}
