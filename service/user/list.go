package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

func (srv *Service) ProcessListResults(c controller.MContext, result interface{}) interface{} {
	return control.SerializeListUsersReply(
		types.CodeOK,
		control.PackSerializeListUserReply(result.([]model.User)))
}
