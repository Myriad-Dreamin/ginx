package objectservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

func (svc *Service) fillPutFields(
	c controller.MContext, object *model.Object,
	req *control.PutObjectRequest) (fields []string) {
	return nil
}
