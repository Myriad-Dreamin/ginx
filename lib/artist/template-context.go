package artist

import "fmt"

type ObjectType int

const (
	ObjectTypeRequest ObjectType = iota
	ObjectTypeReply
	ObjectTypeObject
)

type TemplateContext interface {
	CurrentObjectType() ObjectType
	SetObjectType(ObjectType)
	Clone() TemplateContext

	AppendPackage(pkgPath string)
	MergePackages(pks PackageSet)
	GetPackages() PackageSet

	GetService() ServiceDescription
	GetCategories() []CategoryDescription
	PushCategory(cat CategoryDescription)
	PopCategory() CategoryDescription

	SetMeta(interface{})
	GetMeta() interface{}
}

type TemplateContextImpl struct {
	packages PackageSet

	objType    ObjectType
	svc        ServiceDescription
	categories []CategoryDescription

	meta interface{}
}

func (t *TemplateContextImpl) GetPackages() PackageSet {
	return t.packages
}

func (t *TemplateContextImpl) MergePackages(pks PackageSet) {
	t.packages = inplaceMergePackage(t.packages, pks)
}

func (t *TemplateContextImpl) SetObjectType(oT ObjectType) {
	t.objType = oT
}

func (t *TemplateContextImpl) CurrentObjectType() ObjectType {
	return t.objType
}

func (t *TemplateContextImpl) Clone() TemplateContext {
	return &TemplateContextImpl{
		packages:   t.packages,
		svc:        t.svc,
		categories: t.categories,
		meta:       t.meta,
	}
}

func (t *TemplateContextImpl) PushCategory(cat CategoryDescription) {
	t.categories = append(t.categories, cat)
	return
}

func (t *TemplateContextImpl) PopCategory() (cat CategoryDescription) {
	t.categories, cat = t.categories[:len(t.categories)-1], t.categories[len(t.categories)-1]
	return
}

func (t *TemplateContextImpl) AppendPackage(pkgPath string) {
	if pkgPath == "time" {
		fmt.Println("...")
	}
	t.packages[pkgPath] = true
}

func (t *TemplateContextImpl) GetService() ServiceDescription {
	return t.svc
}

func (t *TemplateContextImpl) GetCategories() []CategoryDescription {
	return t.categories
}

func (t *TemplateContextImpl) SetMeta(meta interface{}) {
	t.meta = meta
}

func (t *TemplateContextImpl) GetMeta() interface{} {
	return t.meta
}

func defaultMeta() interface{} { return nil }

var metaFac = defaultMeta
var newTemplateContext = defaultNewTemplateContext

func SetDefaultMetaFactory(_metaFac func() interface{}) {
	metaFac = _metaFac
}

func SetDefaultNewTemplateContext(_newTemplateContext func(svc ServiceDescription) TemplateContext) {
	newTemplateContext = _newTemplateContext
}

func defaultNewTemplateContext(svc ServiceDescription) TemplateContext {
	return &TemplateContextImpl{
		packages: make(PackageSet),
		svc:      svc,
		meta:     metaFac(),
	}
}
