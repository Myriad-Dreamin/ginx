package auth

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"strconv"
)

type needsEntity struct{}

func (needsEntity) CreateObj(groupID uint) string {
	return "needs:" + strconv.Itoa(int(groupID))
}

func (needsEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "needs", Action: "r/.*"}
}

func (needsEntity) AddReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, NeedsEntity.CreateObj(aim), "r/.*")
}

func (needsEntity) RemoveReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, NeedsEntity.CreateObj(aim), "r/.*")
}

func (needsEntity) HasReadPolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, NeedsEntity.CreateObj(aim), "r/.*")
}

func (needsEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "needs", Action: "w/.*"}
}

func (needsEntity) AddWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, NeedsEntity.CreateObj(aim), "w/.*")
}

func (needsEntity) RemoveWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, NeedsEntity.CreateObj(aim), "w/.*")
}

func (needsEntity) HasWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, NeedsEntity.CreateObj(aim), "w/.*")
}

func (needsEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "needs", Action: "w/s"}
}

func (needsEntity) AddJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, NeedsEntity.CreateObj(aim), "w/s")
}

func (needsEntity) RemoveJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, NeedsEntity.CreateObj(aim), "w/s")
}

func (needsEntity) HasJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, NeedsEntity.CreateObj(aim), "w/s")
}
