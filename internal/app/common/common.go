package common

import "time"

type Storager interface {
	Register(email, password string) error
	Authenticate(email, password string) error
	StoreSecrets(service, login, password, userEmail string) error
	GetAllSecrets(userEmail string) (AllSecrets, error)
	GeServiceRelatedSecrets(userEmail, service string) (AllSecrets, error)
	EditSecret(userEmail string, secretToEdit SecretToEdit) error
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

type SecretToEdit struct {
	Uuid     string `json:"uuid"`
	Service  string `json:"service"`
	Login    string `json:"login"`
	Password string `json:"password"`
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
