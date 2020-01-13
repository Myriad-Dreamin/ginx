package serial

// ProposingService is the Interface of VirtualService
// Getter of Virtual Service
// Get/Set
//    Faz
//    Models
//    Name
//    FilePath
type ProposingService interface {
	CateOf(faz *Category) ProposingService
	UseModel(models ...*model) ProposingService
	Name(name string) ProposingService
	ToFile(fileName string) ProposingService

	GetFaz() *Category
	GetModels() []*model
	GetName() string
	GetFilePath() string
}

// VirtualService has father and sons, and this is just an abstract service that
// for deriving other class
// Getter of Virtual Service
// Get/Set
//    Faz
//    Models
//    Name
//    FilePath
type VirtualService struct {
	faz      *Category
	models   []*model
	name     string
	filePath string
}

// Getter of Virtual Service

func (v *VirtualService) GetFaz() *Category {
	return v.faz
}

func (v *VirtualService) GetModels() []*model {
	return v.models
}

func (v *VirtualService) GetName() string {
	return v.name
}

func (v *VirtualService) GetFilePath() string {
	return v.filePath
}

// Setter of Virtual Service

func (v *VirtualService) CateOf(faz *Category) ProposingService {
	v.faz = faz
	return v
}

func (v *VirtualService) UseModel(models ...*model) ProposingService {
	v.models = append(v.models, models...)
	return v
}

func (v *VirtualService) Name(name string) ProposingService {
	v.name = name
	return v
}

func (v *VirtualService) ToFile(fileName string) ProposingService {
	v.filePath = fileName
	return v
}
