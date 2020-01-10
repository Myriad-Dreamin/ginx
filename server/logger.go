package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-template/config"
	"go.uber.org/zap/zapcore"
)

func (srv *Server) InstantiateLogger() bool {
	var err error
	srv.Logger, err = logger.NewZapLogger(logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
	if err != nil {
		fmt.Println(err)
		return false
	}
	srv.Module.Provide(config.ModulePath.Global.Logger, srv.Logger)
	return true
}

func (srv *Server) println(a ...interface{}) {
	_, err := fmt.Fprintln(srv.LoggerWriter, a...)
	if err != nil {
		fmt.Println(err)
	}
}

func (srv *Server) printf(format string, a ...interface{}) {
	_, err := fmt.Fprintf(srv.LoggerWriter, format, a...)
	if err != nil {
		fmt.Println(err)
	}
}
