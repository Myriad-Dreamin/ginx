package artist

type TagI = map[string]string

type Type interface {
	String() string
}

type Field interface {
	String() string
}

type ParameterDescription interface {
	GetSource() *source
	GetDTOName() string
	GetType() Type
	GetField() Field
	GetTag() TagI
	GetEmbedObjects() []ObjectDescription
}

type Parameter interface {
	CreateParameterDescription(ctx *Context) ParameterDescription
}

type arrayParam struct {
	p Parameter
}

type arrayParamDescription struct {
	ParameterDescription
	a arrayType
}

type arrayType struct {
	Type
	typeString string
}

func (a arrayType) String() string {
	return "[]" + a.Type.String()
}

func (a *arrayParamDescription) GetType() Type {
	return a.a
}

func (a arrayParam) CreateParameterDescription(ctx *Context) ParameterDescription {
	desc := &arrayParamDescription{
		ParameterDescription: a.p.CreateParameterDescription(ctx),
	}
	desc.a = arrayType{Type: desc.ParameterDescription.GetType()}
	return desc
}

func ArrayParam(param Parameter) Parameter {
	return arrayParam{p: param}
}
