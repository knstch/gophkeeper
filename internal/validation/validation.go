package validation

import (
	"context"
	"regexp"
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/knstch/gophkeeper/internal/app/common"
)

type credentialsToValidate struct {
	email    string
	password string
}

func NewCredentialsToValidate(email, password string) *credentialsToValidate {
	return &credentialsToValidate{
		email:    email,
		password: password,
	}
}

func (credentials credentialsToValidate) ValidateCredentials(ctx context.Context) error {
	if err := validation.ValidateStructWithContext(ctx, &credentials,
		validation.Field(&credentials.email,
			validation.Required.Error(common.ErrFieldIsEmpty),
			validation.By(emailValidation(credentials.email)),
		),
		validation.Field(&credentials.password,
			validation.Required.Error(common.ErrFieldIsEmpty),
			validation.By(passwordValidation(credentials.password)),
		),
	); err != nil {
		return err
	}
	return nil
}

func emailValidation(email string) validation.RuleFunc {
	return func(value interface{}) error {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		r := regexp.MustCompile(emailRegex)
		if !r.MatchString(email) {
			return common.ErrBadEmail
		}
		return nil
	}
}

func passwordValidation(password string) validation.RuleFunc {
	return func(value interface{}) error {
		isValid := func(s string) bool {
			var (
				hasMinLen  = false
				hasUpper   = false
				hasLower   = false
				hasNumber  = false
				hasSpecial = false
			)
			if len(s) >= 8 {
				hasMinLen = true
			}
			for _, char := range s {
				switch {
				case unicode.IsUpper(char):
					hasUpper = true
				case unicode.IsLower(char):
					hasLower = true
				case unicode.IsNumber(char):
					hasNumber = true
				case unicode.IsPunct(char) || unicode.IsSymbol(char):
					hasSpecial = true
				}
			}
			return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
		}

		if !isValid(password) {
			return common.ErrBadPass
		}
		return nil
	}
}
