package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/core-oj/log"
)

func (srv *Server)  InstantiateLogger() bool {
	var err error
	srv.Logger, err = log.NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}