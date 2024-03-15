package handler

import (
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/service/auth"
	"github.com/knstch/gophkeeper/internal/app/service/secret"
)

type Handlers struct {
	AuthService   *auth.AuthService
	SecretService *secret.SecretService
}

func NewHandler(storage common.Storager) *Handlers {
	return &Handlers{
		AuthService:   auth.NewAuthService(storage),
		SecretService: secret.NewSecretService(storage),
	}
}

type Message struct {
	Msg   string `json:"message"`
	Error string `json:"error"`
}
