package common

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

type Storager interface {
	AuthStorager
	SecretsStorager
	TextStorager
	BankStorager
}

type AuthStorager interface {
	Register(email, password string) error
	Authenticate(email, password string) error
}

type SecretsStorager interface {
	StoreSecrets(service, login, password, userEmail, metadata string) error
	GetAllSecrets(userEmail string) (*AllSecrets, error)
	GetServiceRelatedSecrets(userEmail, service string) (*AllSecrets, error)
	EditSecret(userEmail, uuid, service, login, password, metadata string) error
	DeleteSecret(userEmail, uuid string) error
}

type TextStorager interface {
	AddTextData(text, title, userEmail, metadata string) error
	EditTextData(text, title, userEmail, metadata, uuid string) error
	DeleteTextData(userEmail, uuid string) error
	GetAllTexts(userEmail string) (*AllTexts, error)
	GetTitleRelatedText(userEmail, title string) (*AllTexts, error)
}

type BankStorager interface {
	StoreCard(userEmail, bankName, cardNumber, date, holderName, metadata string, cvv int) error
	GetAllCards(userEmail string) (*AllCards, error)
	GetBankRelatedCards(userEmail, bankName string) (*AllCards, error)
	EditBankCard(userEmail, bankName, cardNumber, date, holderName, metadata, uuid string, cvv int) error
	DeleteCard(userEmail, uuid string) error
}

type Secrets struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uuid      string    `json:"uuid"`
	Service   string    `json:"service"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Metadata  string    `json:"metadata"`
}

type AllSecrets struct {
	Secrets []Secrets `json:"secrets"`
}

type Credentials struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Uuid      string
	Email     string
	Password  string
}

type Text struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Metadata  string    `json:"metadata"`
}

type AllTexts struct {
	Texts []Text `json:"texts"`
}

type Card struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Uuid       string    `json:"uuid"`
	Email      string    `json:"email"`
	BankName   string    `json:"bank_name"`
	CardNumber string    `json:"card_number"`
	Date       string    `json:"date"`
	HolderName string    `json:"holder_name"`
	Cvv        int       `json:"cvv"`
	Metadata   string    `json:"metadata"`
}

type AllCards struct {
	Cards []Card `json:"cards"`
}

func RetrieveLogin(c *fiber.Ctx) (string, error) {
	if c.Locals("login") == nil {
		return "", ErrNotLoggedIn
	}

	userEmail := c.Locals("login").(string)

	return userEmail, nil
}
