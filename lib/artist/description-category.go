package artist

type CategoryDescription interface {
	GenTreeNode

	GetName() string
	GetPath() string
	GetMethods() []MethodDescription
}

type categoryDescription struct {
	name          string
	path          string
	subCates      map[string]CategoryDescription
	tmplFactories []FunctionTemplateFactory
	methods       []MethodDescription
	packages      PackageSet
}

func (c2 *categoryDescription) GenerateObjects(ts []FunctionTemplateFactory, c TemplateContext) (objs []ObjectTemplate, funcs []FunctionTemplate) {
	return GenerateObjects(c2, ts, c)
}

func (c2 *categoryDescription) GetTemplateFunctionFactory() []FunctionTemplateFactory {
	return c2.tmplFactories
}

func (c *categoryDescription) GetPath() string {
	return c.path
}

func (c *categoryDescription) GetCategories() (categories []CategoryDescription) {
	for _, x := range c.subCates {
		categories = append(categories, x)
	}
	return
}

func (c *categoryDescription) GetMethods() []MethodDescription {
	return c.methods
}

func (c *categoryDescription) GetName() string {
	return c.name
}

func (c *categoryDescription) GetPackages() PackageSet {
	if c == nil {
		return nil
	}
	pac := clonePackage(c.packages)
	//for _, cate := range c.subCates {
	//	pac = inplaceMergePackage(pac, cate.GetPackages())
	//}
	return pac
}
