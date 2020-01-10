package userservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/model"
	"github.com/Myriad-Dreamin/minimum-template/types"
)

type ListReply struct {
	Code  types.CodeType `json:"code"`
	Users []GetReply     `json:"users"`
}

func UsersToListReply(obj []model.User) (reply *ListReply) {
	reply = new(ListReply)
	reply.Code = types.CodeOK
	reply.Users = make([]GetReply, len(obj))
	for i := range obj {
		reply.Users[i] = UserToGetReply(&obj[i])
	}
	return
}

func (srv *Service) ProcessListResults(c controller.MContext, result interface{}) interface{} {
	return UsersToListReply(result.([]model.User))
}
