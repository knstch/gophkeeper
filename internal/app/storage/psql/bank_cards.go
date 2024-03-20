package psql

import (
	"time"

	"github.com/google/uuid"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (storage *PsqlStorage) StoreCard(userEmail, bankName, cardNumber, date, holderName, metadata string, cvv int) error {
	card := &common.Card{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Uuid:       uuid.New().String(),
		Email:      userEmail,
		BankName:   bankName,
		CardNumber: cardNumber,
		Date:       date,
		HolderName: holderName,
		Cvv:        cvv,
		Metadata:   metadata,
	}

	if err := storage.db.Create(card).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqlStorage) GetAllCards(userEmail string) (*common.AllCards, error) {
	var cards []common.Card

	if err := storage.db.Where("email = ?", userEmail).Find(&cards).Error; err != nil {
		return &common.AllCards{}, err
	}

	return &common.AllCards{
		Cards: cards,
	}, nil
}

func (storage *PsqlStorage) GetBankRelatedCards(userEmail, bankName string) (*common.AllCards, error) {
	var cards []common.Card

	if err := storage.db.Where("email = ? AND bank_name = ?", userEmail, bankName).
		Find(&cards).Error; err != nil {
		return &common.AllCards{}, err
	}

	return &common.AllCards{
		Cards: cards,
	}, nil
}

func (storage *PsqlStorage) EditBankCard(userEmail, bankName, cardNumber,
	date, holderName, metadata, uuid string, cvv int) error {
	var checkCard common.Card

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).Find(&checkCard).Error; err != nil {
		return err
	}

	if checkCard.BankName != bankName {
		checkCard.BankName = bankName
	}
	if checkCard.CardNumber != cardNumber {
		checkCard.CardNumber = cardNumber
	}
	if checkCard.Date != date {
		checkCard.Date = date
	}
	if checkCard.HolderName != holderName {
		checkCard.HolderName = holderName
	}
	if checkCard.Metadata != metadata {
		checkCard.Metadata = metadata
	}
	if checkCard.Cvv != cvv {
		checkCard.Cvv = cvv
	}

	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).Save(&common.Card{
		CreatedAt:  checkCard.CreatedAt,
		UpdatedAt:  time.Now(),
		Uuid:       checkCard.Uuid,
		Email:      checkCard.Email,
		BankName:   checkCard.BankName,
		CardNumber: checkCard.CardNumber,
		Date:       checkCard.Date,
		HolderName: checkCard.HolderName,
		Cvv:        checkCard.Cvv,
		Metadata:   checkCard.Metadata,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (storage *PsqlStorage) DeleteCard(userEmail, uuid string) error {
	if err := storage.db.Where("email = ? AND uuid = ?", userEmail, uuid).Delete(&common.Card{}).Error; err != nil {
		if err != nil {
			return err
		}
	}
	return nil
}
