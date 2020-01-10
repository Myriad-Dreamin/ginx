package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

type PutRequest struct {
}

func (srv *Service) fillPutFields(c controller.MContext, object *model.Object, req *PutRequest) (fields []string) {
	return nil
}
