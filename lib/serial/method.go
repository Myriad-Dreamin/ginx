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
	requests []*RequestObj
	replies  []*ReplyObj
}

type context struct {
	vars map[string]interface{}
	method *Method
	svc ProposingService
	packages map[string]int
}

func (c *context) appendPackage(pkg string) {
	if c.packages == nil {
		c.packages = make(map[string]int)
	}
	c.packages[pkg] = 1
}

func (c *context) set(k string, v interface{}) {
	if c.vars == nil {
		c.vars = make(map[string]interface{})
	}
	c.vars[k] = v
}

func (c *context) get(k string) (v interface{}) {
	if c.vars != nil {
		v, _ = c.vars[k]
	}
	return
}

func (method *Method) create(ctx *context) *methodDescription {
	desc := new(methodDescription)
	ctx.method = method
	desc.method = method.method
	desc.name = method.name
	for _, request := range method.requests {
		desc.requests = append(desc.requests, request.create(ctx))
	}
	for _, reply := range method.replies {
		desc.replies = append(desc.replies, reply.create(ctx))
	}
	return desc
}

func newMethod(method MethodType) *Method {
	return &Method{method:method}
}


