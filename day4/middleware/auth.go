package middleware

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Token struct {
	Token string `json:"token"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type AuthMiddleware struct{}

func GetRSAKeys() ([]byte, []byte, error) {
	pubKey, err := ioutil.ReadFile("id_rsa.pub")

	if err != nil {
		return nil, nil, nil
	}

	privateKey, err := ioutil.ReadFile("id_rsa")

	if err != nil {
		return nil, nil, nil
	}

	return pubKey, privateKey, err

}

func (a *AuthMiddleware) GenerateToken(privateFile []byte) (string, error) {
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateFile)

	if err != nil {
		return "", err
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (a *AuthMiddleware) GetRSAPublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM(publicKey)
}

func (a *AuthMiddleware) JWTMiddlewareHandler(publicKey []byte) echo.MiddlewareFunc {

	key, err := a.GetRSAPublicKey(publicKey)

	if err != nil {
		panic((err))
	}

	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &jwtCustomClaims{},
		SigningMethod: "RS256",
		SigningKey:    key,
		Skipper: func(c echo.Context) bool {

			// Skip middleware if path is equal 'login'
			if c.Request().URL.Path == "/v1/signin" || (c.Request().URL.Path == "/v1/users" && c.Request().Method == "POST") {
				return true
			}

			return false
		},
	})
}
