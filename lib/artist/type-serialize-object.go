package artist

type SerializeObject interface {
	CreateObjectDescription(ctx *Context) ObjectDescription
}

func Object(descriptions ...interface{}) SerializeObject {
	var parameters []Parameter
	var name string
	for i := range descriptions {
		switch desc := descriptions[i].(type) {
		case SerializeObject:
			return desc
		case Parameter:
			parameters = append(parameters, desc)
		case []Parameter:
			parameters = append(parameters, desc...)
		case string:
			name = desc
		}
	}
	return &serializeObject{
		name:   name,
		params: parameters,
	}
}

type serializeObject struct {
	params []Parameter
	name   string
}

func (obj *serializeObject) CreateObjectDescription(ctx *Context) ObjectDescription {
	desc := new(objectDescription)
	for _, param := range obj.params {
		desc.params = append(desc.params, param.CreateParameterDescription(ctx))
	}
	desc.name = obj.name
	if len(desc.name) == 0 {
		if suf := ctx.get("obj_suf"); suf != nil {
			if suf, ok := suf.(string); ok {
				desc.name = ctx.method.GetName() + suf
			}
		}
	}
	//fmt.Println("creating", desc.name)
	//for i := range desc.params {
	//	fmt.Print("    ")
	//	param := desc.params[i]
	//	fmt.Println(param.fieldName, param.typeString, param.tags, "||", param.embedObjects)
	//}
	return desc
}
