package auth

import (
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/pkg/service/jwt"
)

type AuthService struct {
	email    string
	password string
	storage  common.Storager
}

func NewAuthService(storage common.Storager) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (auth *AuthService) SignUp(email, password string) (string, error) {
	if err := auth.storage.Register(email, password); err != nil {
		return "", err
	}

	jwt, err := jwt.BuildJWT(auth.email)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (auth *AuthService) SignIn(email, password string) (string, error) {
	if err := auth.storage.Authenticate(email, password); err != nil {
		return "", err
	}

	jwt, err := jwt.BuildJWT(auth.email)
	if err != nil {
		return "", nil
	}

	return jwt, nil
}
