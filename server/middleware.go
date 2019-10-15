package server

import (
	"github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	"github.com/Myriad-Dreamin/ginx/router"
	ginhelper "github.com/Myriad-Dreamin/ginx/service/gin-helper"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/gin-contrib/cors"
	//"github.com/Myriad-Dreamin/gin-middleware/auth/privileger"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (srv *Server) PrepareMiddleware() bool {
	srv.jwtMW = jwt.NewMiddleWare(func() *jwt.CustomClaims {
		var cc = new(jwt.CustomClaims)
		cc.CustomField = &types.CustomFields{}
		return cc
	}, func(c *gin.Context, cc *jwt.CustomClaims) error {
		c.Set("uid", strconv.FormatInt(cc.CustomField.(*types.CustomFields).UID, 10))
		return nil
	})
	srv.jwtMW.ExpireSecond = 3600
	srv.jwtMW.RefreshSecond = 3600 * 24 * 7

	srv.routerAuthMW = router.NewMiddleware(&router.NopValidator{},
		"user:", "uid", ginhelper.MissID, ginhelper.AuthFailed)

	srv.corsMW = cors.New(cors.Config{
		//AllowAllOrigins: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowOrigins:     []string{"http://127.0.0.1:80", "https://127.0.0.1:80"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: true,
	})

	return true
}
