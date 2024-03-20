package binary

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/validation"
)

type BinaryService struct {
	Uuid        string
	Email       string
	FileName    string
	ContentType string
	Data        []byte
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
	file.Data = c.Body()

	if err := validation.NewFileToValidate(file.FileName, file.ContentType, &file.Data).ValidateFile(c.Context()); err != nil {
		if err != nil {
			return err
		}
	}

	if err := file.storage.StoreBinary(file.FileName, file.ContentType, email, &file.Data); err != nil {
		return err
	}

	return nil
}

func (file *BinaryService) GetBin(c *fiber.Ctx) (*common.File, error) {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return &common.File{}, err
	}

	readyFile, err := file.storage.GetBinaryFile(email, c.Params("name"), c.Params("uuid"))
	if err != nil {
		return &common.File{}, err
	}

	if readyFile == nil {
		return &common.File{}, common.ErrNoDataWereFound
	}

	return readyFile, nil
}

func (file *BinaryService) DeleteBin(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	err = file.storage.DeleteBinaryFile(email, c.Params("uuid"))
	if err != nil {
		return err
	}

	return nil
}

func (file *BinaryService) EditBinName(c *fiber.Ctx) error {
	email, err := common.RetrieveLogin(c)
	if err != nil {
		return err
	}

	file.FileName = c.Params("name")
	file.Uuid = c.Params("uuid")

	if len(file.FileName) > 50 {
		return fmt.Errorf("название файла не может быть длиннее %d символов", 50)
	}

	err = file.storage.EditBinaryName(email, file.Uuid, file.FileName)
	if err != nil {
		return err
	}

	return nil
}
