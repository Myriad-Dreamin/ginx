package artist

type GenTreeNode interface {
	GenerateObjects(ts []FunctionTemplateFactory, c TemplateContext) (objs []ObjectTemplate, funcs []FunctionTemplate)
	GetCategories() []CategoryDescription
	GetTemplateFunctionFactory() []FunctionTemplateFactory
	GetPackages() PackageSet
}

type ServiceDescription interface {
	GenTreeNode
	GetName() string
	GetBase() string
	GetFilePath() string
}

type serviceDescription struct {
	name          string
	base          string
	tmplFactories []FunctionTemplateFactory
	categories    []CategoryDescription
	filePath      string
	//packages   map[string]int
}

func (description serviceDescription) GetPackages() PackageSet {
	return nil
}

func (description serviceDescription) GetTemplateFunctionFactory() []FunctionTemplateFactory {
	return description.tmplFactories
}

func (description serviceDescription) GetName() string {
	return description.name
}

func (description serviceDescription) GetBase() string {
	return description.base
}

func (description serviceDescription) GetCategories() []CategoryDescription {
	return description.categories
}

func (description serviceDescription) GetFilePath() string {
	return description.filePath
}

func GenerateObjects(
	g GenTreeNode, ts []FunctionTemplateFactory, c TemplateContext) (
	objs []ObjectTemplate, funcs []FunctionTemplate) {
	ctx := c.Clone()
	ctx.MergePackages(g.GetPackages())
	tmplFactories := append(g.GetTemplateFunctionFactory(), ts...)

	for _, cat := range g.GetCategories() {
		ctx.PushCategory(cat)

		os, fs := cat.GenerateObjects(tmplFactories, ctx)
		objs = append(objs, os...)
		funcs = append(funcs, fs...)

		for _, method := range cat.GetMethods() {
			ctx.SetObjectType(ObjectTypeRequest)
			for _, req := range method.GetRequests() {

				os, fs := dumpObj(ctx, tmplFactories, req)
				objs = append(objs, os...)
				funcs = append(funcs, fs...)
			}

			ctx.SetObjectType(ObjectTypeReply)
			for _, res := range method.GetReplies() {
				os, fs := dumpObj(ctx, tmplFactories, res)
				objs = append(objs, os...)
				funcs = append(funcs, fs...)
			}
		}
		ctx.PopCategory()
	}
	return
}
func (description serviceDescription) GenerateObjects(ts []FunctionTemplateFactory, c TemplateContext) (objs []ObjectTemplate, funcs []FunctionTemplate) {
	return GenerateObjects(description, ts, c)
}
