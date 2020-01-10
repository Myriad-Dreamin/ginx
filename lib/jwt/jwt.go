package jwt

import (
	"crypto/rsa"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Middleware customizes the jwt-middleware
type Middleware struct {
	SigningAlgorithm             string
	JWTHeaderKey                 string
	JWTHeaderPrefixWithSplitChar string
	SigningKeyString             string
	SigningKey                   []byte

	// Public key file for asymmetric algorithms
	PubKeyFile string

	// Private key
	privKey *rsa.PrivateKey

	// Public key
	pubKey *rsa.PublicKey

	//MaxRefresh   time.Duration
	RefreshSecond int64
	ExpireSecond  int64

	customClaimsFactory CustomClaimsFactory
	validFunction       CustomClaimsValidateFunction
}

// NewMiddleWare return default middleware setting
func NewMiddleWare(
	customClaimsFactory CustomClaimsFactory,
	validFunction CustomClaimsValidateFunction,
) *Middleware {
	return &Middleware{
		// todo: customize signing algorithm
		SigningAlgorithm:             "HS256",
		JWTHeaderKey:                 "Authorization",
		JWTHeaderPrefixWithSplitChar: "Bearer ",
		SigningKeyString:             GetSignKey(),
		SigningKey:                   []byte(GetSignKey()),
		customClaimsFactory:          customClaimsFactory,
		validFunction:                validFunction,
		// MaxRefresh: default zero
	}
}

// UnauthorizedMessage just return code with reason
type UnauthorizedMessage struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// Build return the middleware
func (middleware *Middleware) Build() controller.HandlerFunc {
	return func(c controller.MContext) {

		claims, err := middleware.CheckIfTokenExpire(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedMessage{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		// then make yourself the custom validation
		// e.g. claims.CustomField.IP == req.IP
		if err = middleware.validFunction(c, claims); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedMessage{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		// store the context for stateful session
		c.Set("claims", claims)
	}
}

// GenerateToken with expired time
func (middleware *Middleware) GenerateToken(field interface{}) (string, error) {
	return middleware.CreateToken(CustomClaims{
		CustomField: field,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,
			ExpiresAt: time.Now().Unix() + middleware.ExpireSecond,
			Issuer:    middleware.SigningKeyString,
		},
	})
}

// GenerateToken with expired time
func (middleware *Middleware) GenerateTokenWithRefreshToken(field interface{}) (string, string, error) {
	c := CustomClaims{
		CustomField: field,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,
			ExpiresAt: time.Now().Unix() + middleware.ExpireSecond,
			Issuer:    middleware.SigningKeyString,
		},
	}
	cs, err := middleware.CreateToken(c)
	if err != nil {
		return "", "", nil
	}
	rs, err := middleware.CreateToken(CustomClaims{
		CustomField: field,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,
			ExpiresAt: time.Now().Unix() + middleware.RefreshSecond,
			Issuer:    middleware.SigningKeyString,
		},
		IsRefreshToken: true,
		RefreshTarget:  &c,
	})
	if err != nil {
		return "", "", nil
	}
	return cs, rs, nil
}

// CreateToken generate a token
func (middleware *Middleware) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middleware.SigningKey)
}

// RefreshToken if ok
func (middleware *Middleware) RefreshToken(c controller.MContext) (string, error) {
	claims, err := middleware.CheckIfTokenExpire(c)
	if err != nil {
		return "", err
	}
	if claims.IsRefreshToken {
		claims.RefreshTarget.ExpiresAt = jwt.TimeFunc().Unix() + middleware.ExpireSecond
		return middleware.CreateToken(*claims.RefreshTarget)
	} else {
		return "", ErrInvalidAuthHeader
	}
}

// RefreshToken if ok
func (middleware *Middleware) RefreshTokenFunc(c controller.MContext, operate func(*CustomClaims) error) func(c controller.MContext) (string, error) {
	return func(c controller.MContext) (string, error) {
		claims, err := middleware.CheckIfTokenExpire(c)
		if err != nil {
			return "", err
		}
		if claims.IsRefreshToken {
			claims.RefreshTarget.ExpiresAt = jwt.TimeFunc().Unix() + middleware.RefreshSecond
			err = operate(claims)
			if err != nil {
				return "", err
			}
			return middleware.CreateToken(*claims.RefreshTarget)
		} else {
			return "", ErrInvalidAuthHeader
		}
	}
}

// CheckIfTokenExpire check if token expire
func (middleware *Middleware) CheckIfTokenExpire(c controller.MContext) (*CustomClaims, error) {
	token, err := middleware.ParseToken(c)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return nil, err
		}
	}

	if token == nil {
		return nil, ErrEmptyAuthHeader
	}

	claims := token.Claims.(*CustomClaims)

	if claims.ExpiresAt < jwt.TimeFunc().Unix() {
		return nil, ErrExpiredToken
	}

	return claims, nil
}

func (middleware *Middleware) ParseWithClaims(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, middleware.customClaimsFactory(), middleware.KeyFunc)
}

// ParseToken check and return if token in the context
func (middleware *Middleware) ParseToken(c controller.MContext) (*jwt.Token, error) {
	token, err := middleware.jwtFromHeader(c)
	if err != nil {
		return nil, err
	}

	return middleware.ParseWithClaims(token)
}

func (middleware *Middleware) KeyFunc(t *jwt.Token) (interface{}, error) {
	if jwt.GetSigningMethod(middleware.SigningAlgorithm) != t.Method {
		return nil, ErrInvalidSigningAlgorithm
	}
	if middleware.usingPublicKeyAlgorithm() {
		return middleware.pubKey, nil
	}

	// save token string if valid
	//c.Set("JWT_TOKEN", token)

	return middleware.SigningKey, nil
}

func (middleware *Middleware) jwtFromHeader(c controller.MContext) (string, error) {
	authHeader := c.GetHeader(middleware.JWTHeaderKey)
	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	if !strings.HasPrefix(authHeader, middleware.JWTHeaderPrefixWithSplitChar) {
		return "", ErrInvalidAuthHeader
	}

	return authHeader[len(middleware.JWTHeaderPrefixWithSplitChar):], nil
}

func (middleware *Middleware) usingPublicKeyAlgorithm() bool {
	switch middleware.SigningAlgorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}
