package auth

import (
	"context"
	"encoding/json"

	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/pkg/service/jwt"
	"github.com/knstch/gophkeeper/internal/validation"
)

type AuthService struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	storage  common.Storager
}

func NewAuthService(storage common.Storager) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (auth *AuthService) SignUp(ctx context.Context, body []byte) (string, error) {
	if err := json.Unmarshal(body, auth); err != nil {
		return "", err
	}
	if err := validation.NewCredentialsToValidate(auth.Email, auth.Password).
		ValidateCredentials(ctx); err != nil {
		return "", err
	}

	if err := auth.storage.Register(auth.Email, auth.Password); err != nil {
		return "", err
	}

	jwt, err := jwt.BuildJWT(auth.Email)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (auth *AuthService) SignIn(body []byte) (string, error) {
	if err := json.Unmarshal(body, auth); err != nil {
		return "", err
	}
	if err := auth.storage.Authenticate(auth.Email, auth.Password); err != nil {
		return "", err
	}

	jwt, err := jwt.BuildJWT(auth.Email)
	if err != nil {
		return "", nil
	}

	return jwt, nil
}
