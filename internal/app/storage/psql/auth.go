package psql

import (
	"time"

	"github.com/google/uuid"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (storage *PsqlStorage) Register(email, password string) error {
	user := &common.Credentials{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Uuid:      uuid.New().String(),
		Email:     email,
		Password:  password,
	}

	if err := storage.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqlStorage) Authenticate(email, password string) error {
	user := &common.Credentials{}

	if err := storage.db.Where("email = ? AND password = ?", email, password).
		First(user).Error; err != nil {
		return err
	}

	return nil
}
