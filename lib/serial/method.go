package serial

type MethodType int
const (
	POST MethodType = iota
	GET
	PUT
	DELETE
	OPTION
)


type Method struct {
	method   MethodType
	name     string
	requests []SerializeObjectI
	replies []SerializeObjectI
}

func (method *Method) create(ctx *Context) *methodDescription {
	desc := new(methodDescription)
	ctx.method = method
	desc.method = method.method
	desc.name = method.name
	for _, request := range method.requests {
		desc.requests = append(desc.requests, request.CreateObjectDescription(ctx))
	}
	for _, reply := range method.replies {
		desc.replies = append(desc.replies, reply.CreateObjectDescription(ctx))
	}
	return desc
}

func newMethod(method MethodType) *Method {
	return &Method{method:method}
}


