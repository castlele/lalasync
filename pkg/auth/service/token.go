package service

import (
	"os"
	"time"

	"github.com/castlele/lalasync/pkg/auth/models"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET_KEY = "JWT_SECRET"

func NewToken(userName string) (string, error) {
	claims := models.UserClaims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(getToken())
}

func getToken() []byte {
	return []byte(os.Getenv(JWT_SECRET_KEY))
}
