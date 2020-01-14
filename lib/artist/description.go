package artist

import "fmt"

type pureType struct {
	typeString string
}

func (p pureType) String() string {
	return p.typeString
}

type parameterDescription struct {
	embedObjects []ObjectDescription
	pType        Type
	name         string
	field        Field
	source       *source
	tags         map[string]string
}

func (p *parameterDescription) GetSource() *source {
	return p.source
}

func (p *parameterDescription) GetDTOName() string {
	return p.name
}

func (p *parameterDescription) GetField() Field {
	return p.field
}

func (p *parameterDescription) GetTag() TagI {
	return p.tags
}

func (p *parameterDescription) GetEmbedObjects() []ObjectDescription {
	return p.embedObjects
}

func (p *parameterDescription) GetType() Type {
	return p.pType
}

func genStructFields(descriptions []ParameterDescription) (res string) {
	for i := range descriptions {
		desc := descriptions[i]
		if len(res) != 0 {
			res += "\n"
		}
		res += "    " + desc.GetField().String() + " " + desc.GetType().String() + genTag(desc.GetTag())
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
