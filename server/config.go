package server

import "github.com/Myriad-Dreamin/minimum-template/config"

func (srv *Server) LoadConfig(cfgPath string) bool {
	srv.Cfg = config.Default()
	err := config.Load(srv.Cfg, cfgPath)
	if err != nil {
		srv.Logger.Debug("parse config error", "error", err)
		return false
	}
	srv.Module.Provide(config.ModulePath.Global.Configuration, srv.Cfg)
	return true
}
func (srv *Server) UseDefaultConfig() bool {
	srv.Cfg = config.Default()
	srv.Module.Provide(config.ModulePath.Global.Configuration, srv.Cfg)
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
