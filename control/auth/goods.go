package auth

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/casbin/casbin/v2"
	"strconv"
)

type Enforcer = casbin.SyncedEnforcer

type goodsEntity struct{}

func (goodsEntity) CreateObj(groupID uint) string {
	return "goods:" + strconv.Itoa(int(groupID))
}

func (goodsEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "goods", Action: "r/.*"}
}

func (goodsEntity) AddReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, GoodsEntity.CreateObj(aim), "r/.*")
}

func (goodsEntity) RemoveReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, GoodsEntity.CreateObj(aim), "r/.*")
}

func (goodsEntity) HasReadPolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, GoodsEntity.CreateObj(aim), "r/.*")
}

func (goodsEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "goods", Action: "w/.*"}
}

func (goodsEntity) AddWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, GoodsEntity.CreateObj(aim), "w/.*")
}

func (goodsEntity) RemoveWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, GoodsEntity.CreateObj(aim), "w/.*")
}

func (goodsEntity) HasWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, GoodsEntity.CreateObj(aim), "w/.*")
}

func (goodsEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "goods", Action: "w/s"}
}

func (goodsEntity) AddJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, GoodsEntity.CreateObj(aim), "w/s")
}

func (goodsEntity) RemoveJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, GoodsEntity.CreateObj(aim), "w/s")
}

func (goodsEntity) HasJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, GoodsEntity.CreateObj(aim), "w/s")
}
