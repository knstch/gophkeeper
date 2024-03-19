package handler

import (
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/service/auth"
	"github.com/knstch/gophkeeper/internal/app/service/bank"
	"github.com/knstch/gophkeeper/internal/app/service/secret"
	"github.com/knstch/gophkeeper/internal/app/service/text"
)

type Handlers struct {
	AuthService   *auth.AuthService
	SecretService *secret.SecretService
	TextService   *text.TextService
	BankService   *bank.BankService
}

func NewHandler(storage common.Storager) *Handlers {
	return &Handlers{
		AuthService:   auth.NewAuthService(storage),
		SecretService: secret.NewSecretService(storage),
		TextService:   text.NewTextService(storage),
		BankService:   bank.NewBankService(storage),
	}
}

type Message struct {
	Msg string `json:"message"`
}

type Err struct {
	Error string `json:"error"`
}
