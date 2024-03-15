package common

import "time"

type Storager interface {
	AuthStorager
	SecretsStorager
}

type AuthStorager interface {
	Register(email, password string) error
	Authenticate(email, password string) error
}

type SecretsStorager interface {
	StoreSecrets(service, login, password, userEmail string) error
	GetAllSecrets(userEmail string) (AllSecrets, error)
	GeServiceRelatedSecrets(userEmail, service string) (AllSecrets, error)
	EditSecret(userEmail, uuid, service, login, password string) error
	DeleteSecret(userEmail, uuid string) error
}

type Secrets struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uuid      string    `json:"uuid"`
	Service   string    `json:"service"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
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

type TextData struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
}

type AllTexts struct {
	Texts []TextData `json:"texts"`
}
