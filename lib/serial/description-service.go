package serial

type serviceDescription struct {
	name       string
	base string
	categories []*categoryDescription
	filePath   string
	//packages   map[string]int
}

func (description serviceDescription) generateObjects() (result string) {
	for _, cat := range description.categories {
		for _, method := range cat.methods {
			for _, req := range method.requests {
				result = dumpObj(result, req)
			}

			for _, res := range method.replies {
				result = dumpObj(result, res)
			}
		}
	}
	return
}

