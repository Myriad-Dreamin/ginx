package splayer

import "github.com/casbin/casbin/v2"

type Enforcer = casbin.SyncedEnforcer

func (s *Provider) Enforcer() *Enforcer {
	return s.enforcer
}
