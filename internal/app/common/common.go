package common

type Storager interface {
	Register(email, password string) error
	Authenticate(email, password string) error
}
