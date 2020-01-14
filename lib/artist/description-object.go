package artist

type ObjectDescription interface {
	GenObjectTemplate() ObjectTemplate
	GetType() Type
	GetEmbedObject() []ObjectDescription
}

type objectDescription struct {
	name   string
	params []ParameterDescription
}

func (desc objectDescription) GetType() Type {
	return pureType{typeString: desc.name}
}

func (desc objectDescription) GetEmbedObject() (dx []ObjectDescription) {
	for _, param := range desc.params {
		dx = append(dx, param.GetEmbedObjects()...)
	}
	return dx
}

func (desc objectDescription) GenObjectTemplate() ObjectTemplate {
	xps := desc.genXParams()
	return &ObjectTemplateImpl{
		// type desc.name struct {
		Name: desc.name, TType: TemplateTypeStruct,
		Fields: genStructFields(desc.params, xps),
		Xps:    xps,
	}
}

func (desc objectDescription) genXParams() (params []*XParam) {
	//desc.params
	for _, param := range desc.params {
		source := param.GetSource()
		if source != nil {
			params = append(params, &XParam{
				name:   source.paramName(),
				typeOf: source.faz.String(),
				source: param,
			})
		} else {
			params = append(params, &XParam{
				name:   "_" + fromSnakeToSmallCamel(param.GetDTOName()),
				typeOf: param.GetType().String(),
				source: param,
			})
		}
	}
	return
}
