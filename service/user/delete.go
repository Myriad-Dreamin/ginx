package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/model"
)

func (srv *Service) deleteHook(c controller.MContext, object *model.User) (canDelete bool) {
	return true
}
