package jwt

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims records authorization information
type CustomClaims struct {
	jwt.StandardClaims
	IsRefreshToken bool
	RefreshTarget  *CustomClaims
	CustomField    interface{}
}

// CustomClaimsFactory is used to generate custom claims for convenient injected fields
type CustomClaimsFactory func() *CustomClaims

type CustomClaimsValidateFunction func(controller.MContext, *CustomClaims) error
