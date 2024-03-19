package text

import (
	"encoding/json"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/validation"
)

type TextService struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Metadata string `json:"metadata"`
	Uuid     string `json:"uuid"`
	storage  common.Storager
}

func NewTextService(storage common.Storager) *TextService {
	return &TextService{
		storage: storage,
	}
}

func (texts *TextService) StoreTexts(c *fiber.Ctx) error {
	if err := json.Unmarshal(c.Body(), &texts); err != nil {
		return err
	}

	if err := validation.NewTextsToValidate(texts.Title, texts.Text, texts.Metadata).
		ValidateText(c.Context()); err != nil {
		return err
	}

	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err = texts.storage.AddTextData(texts.Text, texts.Title, email, texts.Metadata); err != nil {
		return err
	}

	return nil
}

func (texts *TextService) GetAllTexts(c *fiber.Ctx) (common.AllTexts, error) {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return common.AllTexts{}, err
	}

	allTexts, err := texts.storage.GetAllTexts(email)
	if err != nil {
		return common.AllTexts{}, nil
	}

	return *allTexts, nil
}

func (texts *TextService) GetTextByTitle(c *fiber.Ctx) (common.AllTexts, error) {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return common.AllTexts{}, err
	}

	allTexts, err := texts.storage.GetTitleRelatedText(email, c.Params("title"))
	if err != nil {
		return common.AllTexts{}, err
	}

	if len(allTexts.Texts) == 0 {
		return common.AllTexts{}, common.ErroNoDataWereFound
	}

	return *allTexts, nil
}

func (texts *TextService) DeleteText(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err = texts.storage.DeleteTextData(email, c.Params("uuid")); err != nil {
		return err
	}

	return nil
}

func (texts *TextService) EditText(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(c.Body(), &texts); err != nil {
		return err
	}

	if err = texts.storage.EditTextData(texts.Text, texts.Title, email,
		texts.Uuid, texts.Metadata); err != nil {
		return err
	}

	return nil
}
