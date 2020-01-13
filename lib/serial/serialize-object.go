package serial

type SerializeObject struct {
	params []Parameter
	name   string
}

type requestSerializeObjectI struct {
	s SerializeObjectI
}

func (r requestSerializeObjectI) CreateObjectDescription(ctx *Context) ObjectDescription {
	ctx.set("obj_suf", "Request")
	return r.s.CreateObjectDescription(ctx)
}

type replySerializeObjectI struct {
	s SerializeObjectI
}

func (r replySerializeObjectI) CreateObjectDescription(ctx *Context) ObjectDescription {
	ctx.set("obj_suf", "Reply")
	return r.s.CreateObjectDescription(ctx)
}

type RequestObj = requestSerializeObjectI
type ReplyObj = replySerializeObjectI

func (obj *SerializeObject) CreateObjectDescription(ctx *Context) ObjectDescription {
	desc := new(objectDescription)
	for _, param := range obj.params {
		desc.params = append(desc.params, param.Create(ctx))
	}
	desc.name = obj.name
	if len(desc.name) == 0 {
		if suf := ctx.get("obj_suf"); suf != nil {
			if suf, ok := suf.(string); ok {
				desc.name = ctx.method.name + suf
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

func Request(descriptions ...interface{}) RequestObj {

	return requestSerializeObjectI{s: Object(descriptions...)}
}

func Reply(descriptions ...interface{}) ReplyObj {
	return replySerializeObjectI{s: Object(descriptions...)}
}

func Object(descriptions ...interface{}) SerializeObjectI {
	var parameters []Parameter
	var name string
	for i := range descriptions {
		switch desc := descriptions[i].(type) {
		case SerializeObjectI:
			return desc
		case Parameter:
			parameters = append(parameters, desc)
		case string:
			name = desc
		}
	}
	return &SerializeObject{
		name:   name,
		params: parameters,
	}
}
