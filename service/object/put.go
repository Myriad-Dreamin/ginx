package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/gin-gonic/gin"
)

type PutRequest struct {
}

func (srv *Service) fillPutFields(c *gin.Context, object *model.Object, req *PutRequest) (fields []string) {
	return nil
}
