package main

import (
	"flag"
	"github.com/Myriad-Dreamin/ginx/server"
	_ "net/http/pprof"
)

var (
	port = flag.String("port", ":23336", "serve on port")
	isDebug = flag.Bool("debug", false, "serve with debug mode")
)

func main() {
	srv := server.New("./config")
	if srv == nil {
		return
	}

	// srv.Inject(myPlugins...)

	if *isDebug {
		srv.ServeWithPProf(*port)
	} else {
		srv.Serve(*port)
	}


}
