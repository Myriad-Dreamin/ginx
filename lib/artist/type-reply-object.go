package artist

func Reply(descriptions ...interface{}) ReplyObj {
	return replySerializeObjectI{s: Object(descriptions...)}
}

type ReplyObj = replySerializeObjectI
type replySerializeObjectI struct {
	s SerializeObject
}

func (r replySerializeObjectI) CreateObjectDescription(ctx *Context) ObjectDescription {
	ctx.set("obj_suf", "Reply")
	return r.s.CreateObjectDescription(ctx)
}
