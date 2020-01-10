package auth

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"strconv"
)

type userEntity struct{}

func (userEntity) CreateObj(userID uint) string {
	return "user:" + strconv.Itoa(int(userID))
}

func (userEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "user", Action: "r/.*"}
}

func (userEntity) AddReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, UserEntity.CreateObj(aim), "r/.*")
}

func (userEntity) RemoveReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, UserEntity.CreateObj(aim), "r/.*")
}

func (userEntity) HasReadPolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, UserEntity.CreateObj(aim), "r/.*")
}

func (userEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "user", Action: "w/.*"}
}

func (userEntity) AddWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, UserEntity.CreateObj(aim), "w/.*")
}

func (userEntity) RemoveWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, UserEntity.CreateObj(aim), "w/.*")
}

func (userEntity) HasWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, UserEntity.CreateObj(aim), "w/.*")
}

func (userEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "user", Action: "w/s"}
}

func (userEntity) AddJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, UserEntity.CreateObj(aim), "w/s")
}

func (userEntity) RemoveJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, UserEntity.CreateObj(aim), "w/s")
}

func (userEntity) HasJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, UserEntity.CreateObj(aim), "w/s")
}
