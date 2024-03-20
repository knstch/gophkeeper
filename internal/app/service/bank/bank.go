package bank

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/validation"
)

type BankService struct {
	Uuid       string `json:"uuid"`
	Email      string `json:"email"`
	BankName   string `json:"bank_name"`
	CardNumber string `json:"card_number"`
	Date       string `json:"date"`
	HolderName string `json:"holder_name"`
	Cvv        int    `json:"cvv"`
	Metadata   string `json:"metadata"`
	storage    common.Storager
}

func NewBankService(storage common.Storager) *BankService {
	return &BankService{
		storage: storage,
	}
}

func (bank *BankService) StoreCard(c *fiber.Ctx) error {
	if err := json.Unmarshal(c.Body(), &bank); err != nil {
		return err
	}

	if err := validation.NewCardToValidate(bank.BankName, bank.CardNumber, bank.Date, bank.HolderName, bank.Metadata, bank.Cvv).
		ValidateCard(c.Context()); err != nil {
		return err
	}

	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err := bank.storage.StoreCard(email, bank.BankName, bank.CardNumber, bank.Date, bank.HolderName, bank.Metadata, bank.Cvv); err != nil {
		return err
	}

	return nil
}

func (bank *BankService) GetAllCards(c *fiber.Ctx) (*common.AllCards, error) {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return &common.AllCards{}, err
	}

	cards, err := bank.storage.GetAllCards(email)
	if err != nil {
		return &common.AllCards{}, err
	}

	if len(cards.Cards) == 0 {
		return &common.AllCards{}, common.ErrNoDataWereFound
	}

	return cards, nil
}

func (bank *BankService) GetCardsByBankName(c *fiber.Ctx) (*common.AllCards, error) {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return &common.AllCards{}, err
	}

	cards, err := bank.storage.GetBankRelatedCards(email, c.Params("bank"))
	if err != nil {
		return &common.AllCards{}, nil
	}

	if len(cards.Cards) == 0 {
		return &common.AllCards{}, common.ErrNoDataWereFound
	}

	return cards, nil
}

func (bank *BankService) EditBankCard(c *fiber.Ctx) error {
	if err := json.Unmarshal(c.Body(), &bank); err != nil {
		return err
	}

	if err := validation.NewCardToValidate(bank.BankName, bank.CardNumber, bank.Date, bank.HolderName, bank.Metadata, bank.Cvv).
		ValidateCard(c.Context()); err != nil {
		return err
	}

	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err := bank.storage.EditBankCard(email, bank.BankName, bank.CardNumber, bank.Date, bank.HolderName, bank.Metadata, bank.Uuid, bank.Cvv); err != nil {
		return nil
	}

	return nil
}

func (bank *BankService) DeleteBankCard(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err := bank.storage.DeleteCard(email, c.Params("uuid")); err != nil {
		return err
	}

	return nil
}
