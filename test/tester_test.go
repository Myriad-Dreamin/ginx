package tests

import (
	doc_gen "github.com/Myriad-Dreamin/minimum-lib/generate/doc-gen"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/Myriad-Dreamin/minimum-template/server"
	"github.com/Myriad-Dreamin/minimum-template/test/tester"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	sugar.WithFile(func(logFile *os.File) {
		var options = []server.Option{
			server.OptionRouterLoggerWriter{
				Writer: logFile,
			},
		}
		srv = tester.StartTester(options)
		srv.PrintRequest(true)
		srv.CollectResults(true)
		srv.MainM(m)
		srv.DropMock()
		err := doc_gen.FromGinResults(&doc_gen.GinInfo{
			Result:             srv.DumpResults(),
			Host:               "127.0.0.1",
			ControllerProvider: srv.ServiceProvider,
		})
		if err != nil {
			panic(err)
		}
	}, "test.log")
}
