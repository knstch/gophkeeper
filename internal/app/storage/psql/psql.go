package psql

import (
	"fmt"
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
		fmt.Println(err)
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

func (storage *PsqlStorage) EditSecret(userEmail string, secretToEdit common.SecretToEdit) error {
	var checkSecret common.Secrets

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, secretToEdit.Uuid).
		First(&checkSecret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErroNoDataWereFound
		}
		return err
	}

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, secretToEdit.Uuid).
		Save(&common.Secrets{
			Uuid:      checkSecret.Uuid,
			Email:     userEmail,
			UpdatedAt: time.Now(),
			Service:   secretToEdit.Service,
			Login:     secretToEdit.Login,
			Password:  secretToEdit.Password,
		}).Error; err != nil {
		return err
	}
	return nil
}
