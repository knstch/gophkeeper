package validation

import (
	"context"
	"fmt"
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
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.By(emailValidation(credentials.email)),
		),
		validation.Field(&credentials.password,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
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

type secretToValidate struct {
	service  string
	login    string
	password string
	metadata string
}

func NewSecretsToValidate(service, login, password, metadata string) *secretToValidate {
	return &secretToValidate{
		service:  service,
		login:    login,
		password: password,
		metadata: metadata,
	}
}

func (secrets secretToValidate) ValidateSecrets(ctx context.Context) error {
	if err := validation.ValidateStructWithContext(ctx, &secrets,
		validation.Field(&secrets.service,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(1, 255).Error(fmt.Sprintf("значение не может быть больше %d символов", 255)),
		),
		validation.Field(&secrets.login,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(1, 255).Error(fmt.Sprintf("значение не может быть больше %d символов", 255)),
		),
		validation.Field(&secrets.password,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(1, 255).Error(fmt.Sprintf("значение не может быть больше %d символов", 255)),
		),
		validation.Field(&secrets.metadata,
			validation.RuneLength(0, 1000).Error(fmt.Sprintf("значение не может быть больше %d символов", 1000)),
		),
	); err != nil {
		return err
	}
	return nil
}

type textToValidate struct {
	title    string
	text     string
	metadata string
}

func NewTextsToValidate(title, text, metadata string) *textToValidate {
	return &textToValidate{
		title:    title,
		text:     text,
		metadata: metadata,
	}
}

func (text textToValidate) ValidateText(ctx context.Context) error {
	if err := validation.ValidateStructWithContext(ctx, &text,
		validation.Field(&text.title,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(1, 255).Error(fmt.Sprintf("значение не может быть больше %d символов", 255)),
		),
		validation.Field(&text.text,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(1, 65535).Error(fmt.Sprintf("значение не может быть больше %d символов", 65535)),
		),
		validation.Field(&text.metadata,
			validation.RuneLength(0, 1000).Error(fmt.Sprintf("значение не может быть больше %d символов", 1000)),
		),
	); err != nil {
		return err
	}
	return nil
}

type cardToValidate struct {
	BankName   string
	CardNumber string
	Date       string
	HolderName string
	Cvv        int
	Metadata   string
}

func NewCardToValidate(BankName, CardNumber, Date, HolderName, Metadata string, Cvv int) *cardToValidate {
	return &cardToValidate{
		BankName:   BankName,
		CardNumber: CardNumber,
		Date:       Date,
		HolderName: HolderName,
		Metadata:   Metadata,
		Cvv:        Cvv,
	}
}

func (card cardToValidate) ValidateCard(ctx context.Context) error {
	validation.ValidateStructWithContext(ctx, &card,
		validation.Field(&card.BankName,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(1, 35).Error(fmt.Sprintf("название банка не может быть больше %d символов", 35))),
		validation.Field(&card.CardNumber,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(16, 16).Error("номер карты должен быть 16 символов")),
		validation.Field(&card.Cvv,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.RuneLength(3, 3).Error("cvv может быть только 3 символа")),
		validation.Field(&card.Date,
			validation.Required.Error(common.ErrFieldIsEmpty.Error()),
			validation.By(cardDateValidation(card.Date))))
	return nil
}

func cardDateValidation(date string) validation.RuleFunc {
	return func(value interface{}) error {
		dateRegex := `^[0-9]+/[0-9]$`
		r := regexp.MustCompile(dateRegex)
		if !r.MatchString(date) {
			return common.ErrBadCardDate
		}
		return nil
	}
}
