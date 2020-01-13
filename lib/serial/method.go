package serial

type MethodType int

const (
	POST MethodType = iota
	GET
	PATCH
	HEAD
	PUT
	DELETE
	OPTION
	CONNECT
	TRACE
)

type Method interface {
	CreateMethodDescription(ctx *Context) *methodDescription
}


type method struct {
	methodType MethodType
	name       string
	requests   []SerializeObjectI
	replies    []SerializeObjectI
}

func newMethod(methodType MethodType) *method {
	return &method{methodType: methodType}
}

func (method *method) CreateMethodDescription(ctx *Context) *methodDescription {
	desc := new(methodDescription)
	ctx.method = method
	desc.method = method.methodType
	desc.name = method.name
	for _, request := range method.requests {
		desc.requests = append(desc.requests, request.CreateObjectDescription(ctx))
	}
	for _, reply := range method.replies {
		desc.replies = append(desc.replies, reply.CreateObjectDescription(ctx))
	}
	return desc
}
