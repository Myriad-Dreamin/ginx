package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	mgin "github.com/Myriad-Dreamin/minimum-lib/lib/gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var jwtMW *Middleware
var router *gin.Engine

func TestMain(m *testing.M) {
	jwtMW = NewMiddleWare(func() *CustomClaims {
		var cc = new(CustomClaims)
		return cc
	}, func(c controller.MContext, cc *CustomClaims) error {
		return nil
	})
	jwtMW.ExpireSecond = 1
	jwtMW.RefreshSecond = 3

	router = gin.Default()

	router.GET("/", func(c *gin.Context) {
		a, b, e := jwtMW.GenerateTokenWithRefreshToken(nil)
		if e != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, e)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token":         a,
			"refresh_token": b,
		})
	})

	router.GET("/refresh", func(c *gin.Context) {
		newToken, err := jwtMW.RefreshToken(mgin.Context{Context: c})
		if err != nil {
			if ErrExpiredToken == err {
				_ = c.AbortWithError(http.StatusUnauthorized, err)
				return
			}

			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": newToken,
		})
	})
	authRouter := router.Group("/")
	authRouter.Use(mgin.ToGinHandler(jwtMW.Build()))
	authRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"pong": ""})
	})

	m.Run()
}

func TestMiddleware_GenerateToken(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
		return
	}
	router.ServeHTTP(w, req)

	var tokens map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &tokens)

	if w.Code != 200 {
		t.Error("bad code", w.Code, tokens)
		return
	}

	fmt.Println("result", tokens)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header = map[string][]string{
		"Authorization": {"Bearer " + tokens["token"].(string)},
	}
	router.ServeHTTP(w, req)

	var result map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &result)

	if w.Code != 200 {
		t.Error("bad code", w.Code, result)
		return
	}
	fmt.Println("result", result)

	time.Sleep(time.Second * 2)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header = map[string][]string{
		"Authorization": {"Bearer " + tokens["token"].(string)},
	}
	router.ServeHTTP(w, req)

	_ = json.Unmarshal(w.Body.Bytes(), &result)

	if w.Code != 401 {
		t.Error("bad code", result)
		return
	}
	fmt.Println("result", result)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/refresh", nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header = map[string][]string{
		"Authorization": {"Bearer " + tokens["refresh_token"].(string)},
	}
	router.ServeHTTP(w, req)

	_ = json.Unmarshal(w.Body.Bytes(), &result)

	if w.Code != 200 {
		t.Error("bad code", w.Code, result)
		return
	}
	fmt.Println("result", result)
	tokens["token"] = result["token"]

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header = map[string][]string{
		"Authorization": {"Bearer " + tokens["token"].(string)},
	}
	router.ServeHTTP(w, req)

	_ = json.Unmarshal(w.Body.Bytes(), &result)

	if w.Code != 200 {
		t.Error("bad code", w.Code, result)
		return
	}
	fmt.Println("result", result)

	time.Sleep(time.Second * 2)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/refresh", nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header = map[string][]string{
		"Authorization": {"Bearer " + tokens["refresh_token"].(string)},
	}
	router.ServeHTTP(w, req)

	_ = json.Unmarshal(w.Body.Bytes(), &result)

	if w.Code != 401 {
		t.Error("bad code", result)
		return
	}
	fmt.Println("result", result)
}
