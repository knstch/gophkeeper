package psql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqlStorage struct {
	db *gorm.DB
}

func NewPsqlStorage(dsn string) (*PsqlStorage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PsqlStorage{db: db}, nil
}
