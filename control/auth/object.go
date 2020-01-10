package auth

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"strconv"
)

type objectEntity struct{}

func (objectEntity) CreateObj(groupID uint) string {
	return "object:" + strconv.Itoa(int(groupID))
}

func (objectEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "object", Action: "r/.*"}
}

func (objectEntity) AddReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ObjectEntity.CreateObj(aim), "r/.*")
}

func (objectEntity) RemoveReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ObjectEntity.CreateObj(aim), "r/.*")
}

func (objectEntity) HasReadPolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ObjectEntity.CreateObj(aim), "r/.*")
}

func (objectEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "object", Action: "w/.*"}
}

func (objectEntity) AddWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ObjectEntity.CreateObj(aim), "w/.*")
}

func (objectEntity) RemoveWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ObjectEntity.CreateObj(aim), "w/.*")
}

func (objectEntity) HasWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ObjectEntity.CreateObj(aim), "w/.*")
}

func (objectEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "object", Action: "w/s"}
}

func (objectEntity) AddJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ObjectEntity.CreateObj(aim), "w/s")
}

func (objectEntity) RemoveJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ObjectEntity.CreateObj(aim), "w/s")
}

func (objectEntity) HasJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ObjectEntity.CreateObj(aim), "w/s")
}
