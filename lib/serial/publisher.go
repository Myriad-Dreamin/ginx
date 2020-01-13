package serial

import (
	"reflect"
)

type PublishingServices struct {
	rawSvc      []ProposingService
	packageName string
}

func NewService(rawSvc ...ProposingService) *PublishingServices {
	return &PublishingServices{
		rawSvc:      rawSvc,
		packageName: "control",
	}
}

func (c *PublishingServices) AppendService(rawSvc ...ProposingService) *PublishingServices {
	c.rawSvc = append(c.rawSvc, rawSvc...)
	return c
}

func (c *PublishingServices) SetPackageName(packageName string) *PublishingServices {
	c.packageName = packageName
	return c
}

func (c *PublishingServices) GetPackageName() string {
	return c.packageName
}

func (c *PublishingServices) Publish() error {
	return c.Final().Publish()
}

func (c *PublishingServices) Final() (d *PublishedServices) {
	d = new(PublishedServices)
	for _, svc := range c.rawSvc {
		ctx := &Context{
			svc: svc,
		}
		ctx.makeSources()

		desc := &serviceDescription{
			packages: make(map[string]int),
		}
		desc.name = svc.GetName()
		desc.filePath = svc.GetFilePath()
		faz := svc.GetFaz()
		value, svcType := reflect.ValueOf(svc), reflect.TypeOf(svc)
		for svcType.Kind() == reflect.Ptr {
			value, svcType = value.Elem(), svcType.Elem()
		}
		for i := 0; i < value.NumField(); i++ {
			field, fieldType := value.Field(i), svcType.Field(i)
			if fieldType.Type == cateType {
				cate := field.Interface().(*Category)
				if cate != nil {
					desc.categories = append(desc.categories, cate.CreateCategoryDescription(faz, ctx))
				}
			}
		}
		for k := range ctx.packages {
			desc.packages[k] = 1
		}
		d.svc = append(d.svc, desc)
	}
	return d
}
