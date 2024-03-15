package secret

import (
	"encoding/json"

	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/validation"

	fiber "github.com/gofiber/fiber/v2"
)

type SecretService struct {
	Service  string `json:"service"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Uuid     string `json:"uuid"`
	storage  common.Storager
}

func retrieveLogin(c *fiber.Ctx) (string, error) {
	if c.Locals("login") == nil {
		return "", common.ErrNotLoggedIn
	}

	userEmail := c.Locals("login").(string)

	return userEmail, nil
}

func NewSecretService(storage common.Storager) *SecretService {
	return &SecretService{
		storage: storage,
	}
}

func (secret *SecretService) StoreSecret(c *fiber.Ctx) error {
	if err := json.Unmarshal(c.Body(), &secret); err != nil {
		return err
	}

	if err := validation.NewSecretsToValidate(secret.Service, secret.Login, secret.Password).
		ValidateSecrets(c.Context()); err != nil {
		return err
	}

	if c.Locals("login") == nil {
		return common.ErrNotLoggedIn
	}

	userEmail := c.Locals("login").(string)

	if err := secret.storage.StoreSecrets(secret.Service, secret.Login,
		secret.Password, userEmail); err != nil {
		return err
	}
	secret.storage.GetAllSecrets(userEmail)
	return nil
}

func (secret *SecretService) GetAllSecrets(c *fiber.Ctx) (common.AllSecrets, error) {
	email, err := retrieveLogin(c)
	if err != nil {
		return common.AllSecrets{}, err
	}

	allSecrets, err := secret.storage.GetAllSecrets(email)
	if err != nil {
		return common.AllSecrets{}, err
	}

	return allSecrets, nil
}

func (secret *SecretService) GetServiceSecrets(c *fiber.Ctx) (common.AllSecrets, error) {
	email, err := retrieveLogin(c)
	if err != nil {
		return common.AllSecrets{}, err
	}

	allSecrets, err := secret.storage.GeServiceRelatedSecrets(email, c.Params("service"))
	if err != nil {
		return common.AllSecrets{}, err
	}

	return allSecrets, nil
}

func (secret *SecretService) EditServiceSecrets(c *fiber.Ctx) error {
	email, err := retrieveLogin(c)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(c.Body(), &secret); err != nil {
		return err
	}

	if err := validation.NewSecretsToValidate(secret.Service, secret.Login, secret.Password).
		ValidateSecrets(c.Context()); err != nil {
		return err
	}

	if err = secret.storage.EditSecret(email, secret.Uuid, secret.Service, secret.Login, secret.Password); err != nil {
		return err
	}

	return nil
}

func (secret *SecretService) DeleteSecrets(c *fiber.Ctx) error {
	email, err := retrieveLogin(c)
	if err != nil {
		return err
	}

	if err = secret.storage.DeleteSecret(email, c.Params("uuid")); err != nil {
		return err
	}

	return nil
}
