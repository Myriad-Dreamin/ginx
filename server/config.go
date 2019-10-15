package server

import "github.com/Myriad-Dreamin/ginx/config"

func (srv *Server) LoadConfig(cfgPath string) bool {
	srv.cfg = new(config.ServerConfig)
	err := config.Load(srv.cfg, cfgPath)
	if err != nil {
		srv.Logger.Debug("parse config error", "error", err)
		return false
	}
	return true
}

func (srv *Server) FetchConfig(cfg interface{}, cfgPath string) bool {
	err := config.LoadStatic(cfg, cfgPath)
	if err != nil {
		srv.Logger.Debug("parse config error", "error", err)
		return false
	}
	return true
}

