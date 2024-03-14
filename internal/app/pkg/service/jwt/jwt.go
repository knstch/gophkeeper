package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/knstch/gophkeeper/cmd/config"
)

type Claims struct {
	jwt.RegisteredClaims
	Email string
}

func BuildJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, err := token.SignedString([]byte(config.ReadyConfig.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
