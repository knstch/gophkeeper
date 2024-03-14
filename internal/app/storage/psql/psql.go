package psql

import (
	"math/rand"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqURLlStorage struct {
	db *gorm.DB
}

type Credentials struct {
	gorm.Model
	Email    string
	Password string
}

func NewPsqlStorage(dsn string) (*PsqURLlStorage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PsqURLlStorage{db: db}, nil
}

func (storage *PsqURLlStorage) Register(email, password string) error {
	user := &Credentials{
		Email:    email,
		Password: password,
	}
	user.ID = uint(rand.Uint64() % 1e8)
	if err := storage.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqURLlStorage) Authenticate(email, password string) error {
	user := &Credentials{}

	if err := storage.db.Where("email = ? AND password = ?", email, password).
		First(user).Error; err != nil {
		return err
	}

	return nil
}
