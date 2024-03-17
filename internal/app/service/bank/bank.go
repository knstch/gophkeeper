package bank

import (
	"github.com/knstch/gophkeeper/internal/app/common"
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

// func (bank *BankService) StoreCard(c *fiber.Ctx) error {
// 	if err := json.Unmarshal(c.Body(), &bank); err != nil {
// 		return err
// 	}

// }
