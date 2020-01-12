package serial

type Parameter interface{
	Create(ctx *context) *ParameterDescription
}


type arrayParam struct {
	p Parameter
}

func (a arrayParam) Create(ctx *context) *ParameterDescription {
	desc := a.p.Create(ctx)
	desc.TypeString = "[]" + desc.TypeString
	return desc
}

func ArrayParam(param Parameter) Parameter {
	return arrayParam{p:param}
}



