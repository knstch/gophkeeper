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
	Metadata string `json:"metadata"`
	storage  common.Storager
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

	if err := validation.NewSecretsToValidate(secret.Service, secret.Login, secret.Password, secret.Metadata).
		ValidateSecrets(c.Context()); err != nil {
		return err
	}

	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err := secret.storage.StoreSecrets(secret.Service, secret.Login,
		secret.Password, email, secret.Metadata); err != nil {
		return err
	}

	return nil
}

func (secret *SecretService) GetAllSecrets(c *fiber.Ctx) (common.AllSecrets, error) {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return common.AllSecrets{}, err
	}

	allSecrets, err := secret.storage.GetAllSecrets(email)
	if err != nil {
		return common.AllSecrets{}, err
	}

	if len(allSecrets.Secrets) == 0 {
		return common.AllSecrets{}, common.ErrNoDataWereFound
	}

	return *allSecrets, nil
}

func (secret *SecretService) GetSecretsByService(c *fiber.Ctx) (common.AllSecrets, error) {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return common.AllSecrets{}, err
	}

	allSecrets, err := secret.storage.GetServiceRelatedSecrets(email, c.Params("service"))
	if err != nil {
		return common.AllSecrets{}, err
	}

	if len(allSecrets.Secrets) == 0 {
		return common.AllSecrets{}, common.ErrNoDataWereFound
	}

	return *allSecrets, nil
}

func (secret *SecretService) EditServiceSecrets(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(c.Body(), &secret); err != nil {
		return err
	}

	if err := validation.NewSecretsToValidate(secret.Service, secret.Login, secret.Password, secret.Metadata).
		ValidateSecrets(c.Context()); err != nil {
		return err
	}

	if err = secret.storage.EditSecret(email, secret.Uuid, secret.Service, secret.Login, secret.Password, secret.Metadata); err != nil {
		return err
	}

	return nil
}

func (secret *SecretService) DeleteSecrets(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err = secret.storage.DeleteSecret(email, c.Params("uuid")); err != nil {
		return err
	}

	return nil
}
