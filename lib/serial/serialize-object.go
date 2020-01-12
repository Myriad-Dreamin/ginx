package serial

import "fmt"

type SerializeObject struct {
	params []Parameter
	name   string
}

type RequestObj SerializeObject

func (obj *RequestObj) create(ctx *context) *objectDescription {
	ctx.set("obj_suf", "Request")
	return (*SerializeObject)(obj).create(ctx)
}

func (obj *ReplyObj) create(ctx *context) *objectDescription {
	ctx.set("obj_suf", "Reply")
	return (*SerializeObject)(obj).create(ctx)
}

func (obj *SerializeObject) create(ctx *context) *objectDescription {
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
	fmt.Println("creating", desc.name)
	for i := range desc.params {
		fmt.Print("    ")
		param := desc.params[i]
		fmt.Println(param.fieldName, param.TypeString, param.tags, "||", param.embedObjects)
	}
	return desc
}
type ReplyObj SerializeObject

func Request(descriptions ...interface{}) *RequestObj {
	return (*RequestObj)(Object(descriptions...))
}

func Reply(descriptions ...interface{}) *ReplyObj {
	return (*ReplyObj)(Object(descriptions...))
}

func Object(descriptions ...interface{}) *SerializeObject {
	var parameters []Parameter
	var name string
	for i := range descriptions {
		switch desc := descriptions[i].(type) {
		case Parameter:
			parameters = append(parameters, desc)
		case string:
			name = desc
		}
	}
	return &SerializeObject {
		name: name,
		params: parameters,
	}
}
