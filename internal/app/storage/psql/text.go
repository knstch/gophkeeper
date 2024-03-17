package psql

import (
	"time"

	"github.com/google/uuid"
	"github.com/knstch/gophkeeper/internal/app/common"
	"gorm.io/gorm"
)

func (storage *PsqlStorage) AddTextData(text, title, userEmail, metadata string) error {

	var checkText *common.Text

	if err := storage.db.Where("email = ? AND title = ?", &userEmail, &title).
		First(&checkText).Error; err == nil {
		return common.ErrTextDouble
	}

	readyText := &common.Text{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Uuid:      uuid.New().String(),
		Email:     userEmail,
		Title:     title,
		Text:      text,
		Metadata:  metadata,
	}

	if err := storage.db.Create(&readyText).Error; err != nil {
		return err
	}

	return nil
}

func (storage *PsqlStorage) EditTextData(text, title, userEmail, metadata, uuid string) error {
	var checkText common.Text

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).
		First(&checkText).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErroNoDataWereFound
		}
		return err
	}

	if checkText.Text != text {
		checkText.Text = text
	}
	if checkText.Title != title {
		checkText.Title = title
	}
	if checkText.Metadata != metadata {
		checkText.Metadata = metadata
	}

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).Save(&common.Text{
		CreatedAt: checkText.CreatedAt,
		UpdatedAt: time.Now(),
		Uuid:      checkText.Uuid,
		Email:     checkText.Email,
		Title:     checkText.Title,
		Text:      checkText.Text,
		Metadata:  checkText.Metadata,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqlStorage) DeleteTextData(userEmail, uuid string) error {
	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).Delete(&common.Text{}).Error; err != nil {
		if err != nil {
			return err
		}
	}
	return nil
}

func (storage *PsqlStorage) GetAllTexts(userEmail string) (*common.AllTexts, error) {
	var texts []common.Text
	if err := storage.db.Where("email = ?", userEmail).Find(&texts).Error; err != nil {
		return &common.AllTexts{}, err
	}
	return &common.AllTexts{
		Texts: texts,
	}, nil
}

func (storage *PsqlStorage) GetTitleRelatedText(userEmail, title string) (*common.AllTexts, error) {
	var texts []common.Text
	if err := storage.db.Where("email = ? AND service = ?", userEmail, title).Find(&texts).Error; err != nil {
		return &common.AllTexts{}, err
	}
	return &common.AllTexts{
		Texts: texts,
	}, nil
}
