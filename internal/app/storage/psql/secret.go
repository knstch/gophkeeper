package psql

import (
	"time"

	"github.com/google/uuid"
	"github.com/knstch/gophkeeper/internal/app/common"
	"gorm.io/gorm"
)

func (storage *PsqlStorage) StoreSecrets(service, login, password, userEmail, metadata string) error {
	secret := &common.Secrets{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Uuid:      uuid.New().String(),
		Service:   service,
		Login:     login,
		Password:  password,
		Email:     userEmail,
		Metadata:  metadata,
	}

	if err := storage.db.Create(secret).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqlStorage) GetAllSecrets(userEmail string) (*common.AllSecrets, error) {
	var secrets []common.Secrets
	if err := storage.db.Where("email = ?", userEmail).Find(&secrets).Error; err != nil {
		return &common.AllSecrets{}, err
	}
	return &common.AllSecrets{
		Secrets: secrets,
	}, nil
}

func (storage *PsqlStorage) GetServiceRelatedSecrets(userEmail, service string) (*common.AllSecrets, error) {
	var secrets []common.Secrets
	if err := storage.db.Where("email = ? AND service = ?", userEmail, service).Find(&secrets).Error; err != nil {
		return &common.AllSecrets{}, err
	}
	return &common.AllSecrets{
		Secrets: secrets,
	}, nil
}

func (storage *PsqlStorage) EditSecret(userEmail, uuid, service, login, password, metadata string) error {
	var checkSecret common.Secrets

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).
		First(&checkSecret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrNoDataWereFound
		}
		return err
	}

	if checkSecret.Service != service {
		checkSecret.Service = service
	}
	if checkSecret.Login != login {
		checkSecret.Login = login
	}
	if checkSecret.Password != password {
		checkSecret.Password = password
	}
	if checkSecret.Metadata != metadata {
		checkSecret.Metadata = metadata
	}

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).
		Save(&common.Secrets{
			CreatedAt: checkSecret.CreatedAt,
			Uuid:      checkSecret.Uuid,
			Email:     userEmail,
			UpdatedAt: time.Now(),
			Service:   checkSecret.Service,
			Login:     checkSecret.Login,
			Password:  checkSecret.Password,
			Metadata:  checkSecret.Metadata,
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
