//

package auth

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/casbin/casbin/v2"
)

type Enforcer = casbin.SyncedEnforcer

// EntityInterface Interface
type EntityInterface interface {
	CreateObj(groupID uint) string

	ReadEntityInterface
	WriteEntityInterface
	JustSimpleWriteEntityInterface
}

// ReadEntityInterface Interface
type ReadEntityInterface interface {
	Read() controller.Requirement
	AddReadPolicy(e *Enforcer, subject string, aim uint) (bool, error)
	RemoveReadPolicy(e *Enforcer, subject string, aim uint) (bool, error)
	HasReadPolicy(e *Enforcer, subject string, aim uint) bool
}

// WriteEntityInterface Interface
type WriteEntityInterface interface {
	Write() controller.Requirement
	AddWritePolicy(e *Enforcer, subject string, aim uint) (bool, error)
	RemoveWritePolicy(e *Enforcer, subject string, aim uint) (bool, error)
	HasWritePolicy(e *Enforcer, subject string, aim uint) bool
}

// JustSimpleWriteEntityInterface Interface
type JustSimpleWriteEntityInterface interface {
	JustSimpleWrite() controller.Requirement
	AddJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error)
	RemoveJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error)
	HasJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) bool
}
