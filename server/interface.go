//

package server

import (
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"github.com/Myriad-Dreamin/minimum-template/lib/plugin"
	"testing"
)

// MockInterface Interface
type MockInterface interface {
	ReleaseMock()
	DropMock()

	PrintRequest(p bool)
	DumpResults() (res []Res)

	CollectResults(collectResults bool) *Mocker
	NoErr(resp mock.ResponseI) bool

	SetHeader(k, v string)
	UseToken(token string)
	Context(t *testing.T) *MockerContext
	FetchError(resp mock.ResponseI) Error

	Method(method, path string, params ...interface{}) mock.ResponseI
	Get(path string, params ...interface{}) mock.ResponseI
	Connect(path string, params ...interface{}) mock.ResponseI
	Delete(path string, params ...interface{}) mock.ResponseI
	Head(path string, params ...interface{}) mock.ResponseI
	Options(path string, params ...interface{}) mock.ResponseI
	Patch(path string, params ...interface{}) mock.ResponseI
	Post(path string, params ...interface{}) mock.ResponseI
	Put(path string, params ...interface{}) mock.ResponseI
	Trace(path string, params ...interface{}) mock.ResponseI
}

// ServerInterface Interface
type ServerInterface interface {
	Terminate()
	Inject(plugins ...plugin.Plugin) (injectSuccess bool)
	Serve(port string)
	ServeWithPProf(port string)
}
