package artist

type ServiceDescription interface {
	GetName() string
	GetBase() string
	GetCategories() []CategoryDescription
	GetFilePath() string
	GenerateObjects() string
}

type serviceDescription struct {
	name       string
	base       string
	categories []CategoryDescription
	filePath   string
	//packages   map[string]int
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

func (description serviceDescription) GenerateObjects() (result string) {
	for _, cat := range description.categories {
		for _, method := range cat.GetMethods() {
			for _, req := range method.GetRequests() {
				result = dumpObj(result, req)
			}

			for _, res := range method.GetReplies() {
				result = dumpObj(result, res)
			}
		}
	}
	return
}
