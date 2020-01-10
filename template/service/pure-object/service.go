//go:generate package-attach-to-path -generate_register_map
package objectservice

import (
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Service struct {
	cfg    *config.ServerConfig
	logger types.Logger
}

func (svc *Service) ObjectSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (svc *Service, err error) {
	svc = new(Service)
	svc.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	svc.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	return
}
