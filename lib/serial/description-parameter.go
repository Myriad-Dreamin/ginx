package serial

import "fmt"

type ParameterDescription struct {
	embedObjects []ObjectDescription
	typeString   string
	name         string
	fieldName    string
	source       *source
	tags         map[string]string
}

func genStructFields(descriptions []*ParameterDescription) (res string) {
	for i := range descriptions {
		desc := descriptions[i]
		if len(res) != 0 {
			res += "\n"
		}
		res += "    " + desc.fieldName + " " + desc.typeString + genTag(desc.tags)
	}
	return
}

func genTag(tags map[string]string) (res string) {
	res = "`"
	for k, v := range tags {
		res += fmt.Sprintf(`%s:"%s"`, k, v)
	}
	res += "`"
	return res
}
