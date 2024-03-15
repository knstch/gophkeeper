package psql

import (
	"time"

	"github.com/google/uuid"
	"github.com/knstch/gophkeeper/internal/app/common"
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

func (storage *PsqlStorage) StoreSecrets(service, login, password, userEmail string) error {
	secret := &common.Secrets{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Uuid:      uuid.New().String(),
		Service:   service,
		Login:     login,
		Password:  password,
		Email:     userEmail,
	}

	if err := storage.db.Create(secret).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqlStorage) GetAllSecrets(userEmail string) (common.AllSecrets, error) {
	var secrets []common.Secrets
	if err := storage.db.Where("email = ?", userEmail).Find(&secrets).Error; err != nil {
		return common.AllSecrets{}, err
	}
	return common.AllSecrets{
		Secrets: secrets,
	}, nil
}

func (storage *PsqlStorage) GeServiceRelatedSecrets(userEmail, service string) (common.AllSecrets, error) {
	var secrets []common.Secrets
	if err := storage.db.Where("email = ? AND service = ?", userEmail, service).Find(&secrets).Error; err != nil {
		return common.AllSecrets{}, err
	}
	return common.AllSecrets{
		Secrets: secrets,
	}, nil
}

func (storage *PsqlStorage) EditSecret(userEmail, uuid, service, login, password string) error {
	var checkSecret common.Secrets

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).
		First(&checkSecret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErroNoDataWereFound
		}
		return err
	}

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).
		Save(&common.Secrets{
			Uuid:      checkSecret.Uuid,
			Email:     userEmail,
			UpdatedAt: time.Now(),
			Service:   service,
			Login:     login,
			Password:  password,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqlStorage) DeleteSecret(userEmail, uuid string) error {
	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).Delete(&common.Secrets{}).Error; err != nil {
		if err != nil {
			return err
		}
	}
	return nil
}

func (storage *PsqlStorage) AddTextData(text, title, userEmail string) error {
	readyText := &common.TextData{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Uuid:      uuid.New().String(),
		Email:     userEmail,
		Title:     title,
		Text:      text,
	}

	if err := storage.db.Create(readyText).Error; err != nil {
		return err
	}

	return nil
}
