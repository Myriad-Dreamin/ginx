package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Myriad-Dreamin/go-magic-package/instance"
	parser "github.com/Myriad-Dreamin/go-parse-package"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/control"
	"github.com/Myriad-Dreamin/minimum-template/lib/serial"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Myriad-Dreamin/gin-middleware/mock"
	abstract_test "github.com/Myriad-Dreamin/minimum-lib/abstract-test"
	"github.com/Myriad-Dreamin/minimum-lib/mock"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	dblayer "github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	"github.com/stretchr/testify/assert"
)

type Mocker struct {
	*Server
	cancel             func()
	header             map[string]string
	routes             map[string]*Results
	contextHelper      abstract_test.ContextHelperInterface
	shouldPrintRequest bool
	assertNoError      bool
	collectResults     bool
}

type MockerContext struct {
	*Mocker
	*assert.Assertions
}

type Res = mock.GinResultI

func Mock(options ...Option) (srv *Mocker) {
	srv = new(Mocker)
	srv.Server = newServer(options)
	srv.header = make(map[string]string)
	if !(srv.InstantiateLogger() &&
		srv.UseDefaultConfig() &&
		srv.PrepareFileSystem() &&
		srv.MockDatabase()) {
		srv = nil
		return
	}

	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if srv == nil {
			srv.Terminate()
		}
	}()

	if !(srv.PrepareMiddleware() &&
		srv.PrepareService() &&
		srv.BuildRouter()) {
		srv = nil
		return
	}

	if err := srv.Module.Install(srv.RouterProvider); err != nil {
		srv.println("install router provider error", err)
	}
	if err := srv.Module.Install(srv.ModelProvider); err != nil {
		srv.println("install database provider error", err)
	}

	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}
	}()

	srv.HttpEngine.Use(mockw.ContextRecorder())
	control.BuildHttp(srv.Router.Root, srv.HttpEngine)
	srv.Module.Debug(srv.Logger)
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if srv == nil {
			cancel()
		}
	}()

	for _, plg := range srv.plugins {
		go plg.Work(ctx)
	}

	if err := dblayer.GetRawInstance().Ping(); err != nil {
		srv.Logger.Debug("database died", "error", err)
		srv = nil
		return
	}
	srv.cancel = cancel
	srv.contextHelper = &abstract_test.ContextHelper{Logger: log.New(srv.LoggerWriter, "mocker", log.Ldate|log.Ltime|log.Llongfile|log.LstdFlags)}

	routes := srv.Router.Root.Routes()
	srv.routes = make(map[string]*Results)
	for _, route := range routes {
		srv.routes[route.Path+"@"+route.Method] = &Results{
			RouteInfo: route,
			Recs:      nil,
		}
	}

	return
}

type Results struct {
	controller.RouteInfo
	Recs []mock.RecordsI
}

func (r Results) GetMethod() string {
	return r.Method
}

func (r Results) GetPath() string {
	return r.Path
}

func (r Results) GetHandlerFunc() interface{} {
	return r.HandlerFunc
}

func (r Results) GetHandler() string {
	return r.Handler
}

func (r Results) GetRecords() []mock.RecordsI {
	return r.Recs
}

func (mocker *Mocker) PrintRequest(p bool) {
	mocker.shouldPrintRequest = p
}

func (mocker *Mocker) DumpResults() (res []Res) {
	for _, v := range mocker.routes {
		res = append(res, v)
	}
	return
}

func (mocker *Mocker) Context(t *testing.T) *MockerContext {
	m := new(Mocker)
	*m = *mocker
	m.contextHelper = t
	return &MockerContext{
		Mocker:     m,
		Assertions: assert.New(t),
	}
}

func (mocker *Mocker) ReleaseMock() {
	if mocker.cancel != nil {
		mocker.cancel()
		mocker.cancel = nil
	}
}

type Request = http.Request

type rc struct {
	*bytes.Buffer
}

func (rc) Close() error {
	return nil
}

func (mocker *Mocker) mockServe(r *Request, params ...interface{}) (w *mock.Response) {

	w = mock.NewResponse()
	var (
		b           []byte
		err         error
		comment     string = "the request url is " + r.URL.String() + ". "
		abortRecord        = false
	)

	for i := range params {
		switch p := params[i].(type) {
		case mock.Comment:
			comment = comment + string(p)
		case mock.AbortRecord:
			abortRecord = bool(p)
		}
	}

	if !abortRecord && mocker.collectResults {
		if r.Body != nil {
			b, err = ioutil.ReadAll(r.Body)
			_ = r.Body.Close()
			if err != nil {
				mocker.contextHelper.Fatal("read failed", "error", err)
			}
			r.Body = rc{bytes.NewBuffer(b)}
		}
	}

	mocker.HttpEngine.ServeHTTP(w, r)

	if mocker.contextHelper != nil && mocker.assertNoError {
		if !mocker.NoErr(w) {
			mocker.contextHelper.Fatal("stopped by assertion")
		}
	}

	if mocker.shouldPrintRequest {
		mocker.println("Method:", r.Method, "url:", r.URL, "http:", r.Proto)
		mocker.println("Request Header:", r.Header)
		mocker.println("Response Header:", w.Header())
	}

	if !abortRecord && mocker.collectResults {
		c := make([]byte, w.Body().Len())
		copy(c, w.Body().Bytes())
		pattern := w.Header().Get("Gin-Context-Matched-Path-Method")
		w.Header().Del("Gin-Context-Matched-Path-Method")
		if results, ok := mocker.routes[pattern]; ok {
			rec := mock.Records{
				RequestBody:  b,
				ResponseBody: c,
				ResponseCode: w.Code(),
				Comment:      comment,
			}
			rec.RequestHeader = make(http.Header)
			for k, v := range r.Header {
				rec.RequestHeader[k] = v
			}
			rec.ResponseHeader = make(http.Header)
			for k, v := range w.Header() {
				rec.ResponseHeader[k] = v
			}
			results.Recs = append(results.Recs, rec)
		} else {
			mocker.contextHelper.Fatal("matched bad route", pattern)
		}
	}

	return
}

func (mocker *Mocker) report(err error) {
	if mocker.contextHelper != nil {
		mocker.contextHelper.Helper()
		mocker.contextHelper.Error(err)
	} else {
		mocker.Logger.Error("error occurs", "error", err)
	}
}

type emptyBody struct{}

func (body emptyBody) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

var _emptyBody = emptyBody{}

func (mocker *Mocker) Method(method, path string, params ...interface{}) mock.ResponseI {
	var (
		body        io.Reader = _emptyBody
		contentType string
		serveParams []interface{}
		r           *http.Request
	)
	for i := range params {
		switch p := params[i].(type) {
		case string, []byte:
			body = mock.NotStruct(p)
		case mock.Serializable:
			var err error
			body, err = p.Serialize()
			if err != nil {
				mocker.report(err)
				return nil
			}
			contentType = p.ContentType()
		case *url.Values:
			body = strings.NewReader(p.Encode())
		case io.Reader:
			body = p
		case *http.Request:
			r = p
		case http.Request:
			r = &p
		case mock.Comment, mock.AbortRecord:
			serveParams = append(serveParams, p)
		default:
			buf := bytes.NewBuffer(nil)
			body = buf
			if err := json.NewEncoder(buf).Encode(p); err != nil {
				mocker.Logger.Error("encode request to json error", "error", err)
			}
			contentType = "application/json"
		}
	}
	if r == nil {
		var err error
		r, err = http.NewRequest(method, path, body)
		if err != nil {
			mocker.report(err)
			return nil
		}
		r.Header.Set("Content-Type", contentType)
		for k, v := range mocker.header {
			r.Header.Set(k, v)
		}
	}
	return mocker.mockServe(r, serveParams...)
}

func (mocker *Mocker) Get(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodGet, path, params...)
}

func (mocker *Mocker) Connect(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodConnect, path, params...)
}

func (mocker *Mocker) Delete(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodDelete, path, params...)
}

func (mocker *Mocker) Head(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodHead, path, params...)
}

func (mocker *Mocker) Options(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodOptions, path, params...)
}

func (mocker *Mocker) Patch(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodPatch, path, params...)
}

func (mocker *Mocker) Post(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodPost, path, params...)
}

func (mocker *Mocker) Put(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodPut, path, params...)
}

func (mocker *Mocker) Trace(path string, params ...interface{}) mock.ResponseI {
	return mocker.Method(http.MethodTrace, path, params...)
}

func (mocker *Mocker) SetHeader(k, v string) {
	mocker.header[k] = v
}

func (mocker *Mocker) UseToken(token string) {
	mocker.header[mocker.jwtMW.JWTHeaderKey] =
		mocker.jwtMW.JWTHeaderPrefixWithSplitChar + token
}

func (mocker *Mocker) CollectResults(collectResults bool) *Mocker {
	mocker.collectResults = collectResults
	return mocker
}

func (mocker *MockerContext) AssertNoError(noErr bool) *MockerContext {
	mocker.assertNoError = noErr
	return mocker
}

func (mocker *Mocker) NoErr(resp mock.ResponseI) bool {
	if mocker.contextHelper == nil {
		panic("only used in test")
	}
	mocker.contextHelper.Helper()
	if resp.Code() != 200 {
		mocker.contextHelper.Error("resp has bad code ", resp.Code())
		return false
	}
	body := resp.Body()
	var obj serial.ErrorSerializer
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		mocker.contextHelper.Error(err)
		return false
	}
	if len(obj.Error) != 0 || obj.Code != 0 {
		mocker.contextHelper.Errorf("Code, Error (%v, %v)", obj.Code, obj.Error)
		return false
	}
	return true
	//if gjson
}

type Error struct {
	RespCode int
	Code     types.CodeRawType `json:"code"`
	Error    string            `json:"error"`
}

func (mocker *Mocker) FetchError(resp mock.ResponseI) Error {
	if mocker.contextHelper == nil {
		panic("only used in test")
	}
	mocker.contextHelper.Helper()
	var obj Error
	body := resp.Body()
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		mocker.contextHelper.Error(err)
		return obj
	}
	obj.RespCode = resp.Code()
	return obj
	//if gjson
}

func init() {
	parser.SetPackageMapper(instance.Get)
}

func (mocker *Mocker) DropMock() {
	mocker.DropFileSystem()
}
