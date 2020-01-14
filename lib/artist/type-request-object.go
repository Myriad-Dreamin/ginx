package artist

func Request(descriptions ...interface{}) RequestObj {
	return requestSerializeObjectI{s: Object(descriptions...)}
}

type RequestObj = requestSerializeObjectI
type requestSerializeObjectI struct {
	s SerializeObject
}

func (r requestSerializeObjectI) CreateObjectDescription(ctx *Context) ObjectDescription {
	ctx.set("obj_suf", "Request")
	return r.s.CreateObjectDescription(ctx)
}
