package artist

type CategoryDescription interface {
	GetName() string
	GetPath() string
	GetMethods() []MethodDescription
	GetPackages() packageSet
}

type categoryDescription struct {
	name     string
	subCates map[string]CategoryDescription
	path     string
	methods  []MethodDescription
	packages packageSet
}

func (c *categoryDescription) GetPath() string {
	return c.path
}

func (c *categoryDescription) GetMethods() []MethodDescription {
	return c.methods
}

func (c *categoryDescription) GetName() string {
	return c.name
}

func (c *categoryDescription) GetPackages() packageSet {
	if c == nil {
		return nil
	}
	pac := clonePackage(c.packages)
	for _, cate := range c.subCates {
		pac = inplaceMergePackage(pac, cate.GetPackages())
	}
	return pac
}
