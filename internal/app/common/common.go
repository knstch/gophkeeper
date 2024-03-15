package common

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

type Storager interface {
	AuthStorager
	SecretsStorager
	TextStorager
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
	EditTextData(text, title, userEmail, uuid, metadata string) error
	DeleteTextData(userEmail, uuid string) error
	GetAllTexts(userEmail string) (*AllTexts, error)
	GetTitleRelatedText(userEmail, title string) (*AllTexts, error)
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
	DeletedAt time.Time
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

func RetrieveLogin(c *fiber.Ctx) (string, error) {
	if c.Locals("login") == nil {
		return "", ErrNotLoggedIn
	}

	userEmail := c.Locals("login").(string)

	return userEmail, nil
}
