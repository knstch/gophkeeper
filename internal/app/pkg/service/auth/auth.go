package auth

type AuthService struct {
	email    string
	password string
}

func NewAuthService(email, password string) *AuthService {
	return &AuthService{
		email:    email,
		password: password,
	}
}
