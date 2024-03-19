package binary

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

type BinaryService struct {
	Uuid        string
	Email       string
	FileName    string
	ContentType string
	Data        *[]byte
	storage     common.Storager
}

func NewBinaryService(storage common.Storager) *BinaryService {
	return &BinaryService{
		storage: storage,
	}
}

func (file *BinaryService) StoreBin(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	file.FileName = c.Params("name")
	file.ContentType = http.DetectContentType(c.Body())
	file.Email = email
	return nil
}
