package artist

type snakeParam struct {
	n *norm
}

func (s snakeParam) CreateParameterDescription(ctx *Context) ParameterDescription {
	if _, ok := s.n.param.Interface().(SerializeObject); ok {
		panic("embed type has no name")
	} else {
		source := parseSource(ctx, s.n)
		if source != nil {
			s.n.name = fromBigCamelToSnake(source.memberName())
		} else {
			panic("non model value has no name")
		}
	}
	return s.n.CreateParameterDescription(ctx)
}

func newSnake(n *norm) *snakeParam {
	return &snakeParam{n: n}
}

func SnakeParam(descriptions ...interface{}) Parameter {
	return newSnake(createNorm("_snaking", descriptions...))
}
