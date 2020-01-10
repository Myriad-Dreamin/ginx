package base_service

type CRUDObjectToolLite interface {
	FObject
	DObject
	RObject
	UObject
	CObject
}

type CRUDService struct {
	Tool CRUDObjectToolLite
	k    string
	CServiceInterface
	RServiceInterface
	UServiceInterface
	DServiceInterface
}

func NewCRUDService(tool CRUDObjectToolLite, k string) CRUDService {
	return CRUDService{
		Tool:              tool,
		k:                 k,
		CServiceInterface: NewCService(tool, k),
		RServiceInterface: NewRService(tool, k),
		UServiceInterface: NewUService(tool, k),
		DServiceInterface: NewDService(tool, k),
	}
}

