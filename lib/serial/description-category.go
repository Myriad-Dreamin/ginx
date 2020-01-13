package serial

type categoryDescription struct {
	name string
	subCates map[string]*categoryDescription
	path     string
	methods  []*methodDescription
	packages map[string]bool
}

func (c *categoryDescription) getPackages() packageSet {
	if c == nil {
		return nil
	}
	pac := clonePackage(c.packages)
	for _, cate := range c.subCates {
		pac = inplaceMergePackage(pac, cate.getPackages())
	}
	return pac
}


