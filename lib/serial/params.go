package serial

type Parameter interface{
	Create(ctx *Context) *ParameterDescription
}


type arrayParam struct {
	p Parameter
}

func (a arrayParam) Create(ctx *Context) *ParameterDescription {
	desc := a.p.Create(ctx)
	desc.typeString = "[]" + desc.typeString
	return desc
}

func ArrayParam(param Parameter) Parameter {
	return arrayParam{p:param}
}



