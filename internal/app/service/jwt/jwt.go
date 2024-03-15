package jwt

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/knstch/gophkeeper/cmd/config"
	"github.com/knstch/gophkeeper/internal/app/common"
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

func getEmail(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(config.ReadyConfig.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", common.ErrInvalidToken
	}

	return claims.Email, nil
}

func GetEmail(req *http.Request) (string, error) {
	signedCookie, err := req.Cookie("auth")
	if err != nil {
		return "", common.ErrNotLoggedIn
	}

	email, err := getEmail(signedCookie.Value)
	if err != nil {
		return "", err
	}

	return email, nil
}
