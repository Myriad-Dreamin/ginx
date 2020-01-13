package serial

import "path/filepath"

type Category struct {
	path string
	methods []*Method
}

func Ink(_ ...interface{}) *Category {
	return new(Category)
}

func (c *Category) Method(m MethodType, descriptions...interface{}) *Category {
	method := newMethod(m)
	for _, description := range descriptions {
		switch desc := description.(type) {
		case string:
			method.name = desc
		case RequestObj:
			method.requests = append(method.requests, desc)
		case ReplyObj:
			method.requests = append(method.requests, desc)
		}
	}
	c.methods = append(c.methods, method)
	return c
}

func (c *Category) Path(path string) *Category {
	c.path = path
	return c
}

func (c *Category) CreateCategoryDescription(faz *Category, ctx *Context) *categoryDescription {

	desc := new(categoryDescription)
	if faz != nil {
		desc.path = filepath.Join(faz.path, c.path)
	}
	for _, method := range c.methods {
		desc.methods = append(desc.methods, method.create(ctx))
		for k := range ctx.packages {
			ctx.appendPackage(k)
		}
	}
	return desc
}

