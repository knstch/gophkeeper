package handler

import (
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/pkg/service/auth"
)

type Handlers struct {
	AuthService *auth.AuthService
}

func NewHandler(storage common.Storager) *Handlers {
	return &Handlers{
		AuthService: auth.NewAuthService(storage),
	}
}
