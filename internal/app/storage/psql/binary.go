package psql

import (
	"time"

	"github.com/google/uuid"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (storage *PsqlStorage) StoreBinary(fileName, contentType, email string, binaryData *[]byte) error {
	preparedData := &common.File{
		CreatedAt:   time.Now(),
		Uuid:        uuid.New().String(),
		Email:       email,
		FileName:    fileName,
		ContentType: contentType,
		Data:        binaryData,
	}

	if err := storage.db.Create(&preparedData).Error; err != nil {
		return err
	}

	return nil
}

func (storage *PsqlStorage) GetBinaryFile(email, filename string) (*common.File, error) {
	var file common.File

	if err := storage.db.Where("file_name = ? AND email = ?", filename, email).Find(&file).Error; err != nil {
		return &common.File{}, err
	}

	return &file, nil
}

func (storage *PsqlStorage) DeleteBinaryFile(email, uuid string) error {
	var file common.File

	if err := storage.db.Where("email = ? AND uuid = ?", email, uuid).Delete(&file).Error; err != nil {
		return err
	}

	return nil
}
