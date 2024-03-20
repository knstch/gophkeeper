package handler

import (
	"fmt"
	"net/http"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) StoreBinFile() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.BinService.StoreBin(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: "поле не может быть пустым",
				})
			}
			if strings.Contains(err.Error(), "не может быть больше") {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			if strings.Contains(err.Error(), "разрешено добавлять файлы не более") {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: common.ErrFileSize.Error(),
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(Err{
				Error: err.Error(),
			})
		}

		return c.Status(http.StatusAccepted).JSON(Message{
			Msg: "файл успешно сохранен!",
		})
	}
}

func (h *Handlers) GetBinFile() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		file, err := h.BinService.GetBin(c)
		if err != nil {
			if err == common.ErrNoDataWereFound {
				return c.Status(http.StatusNoContent).JSON(Message{
					Msg: "ничего не найдено",
				})
			}
		}

		contentType := selectContentType(file.ContentType)
		if contentType == "" {
			return c.Status(http.StatusInternalServerError).JSON(Err{
				Error: err.Error(),
			})
		}

		c.Set(fiber.HeaderContentDisposition, contentType)
		return c.Status(http.StatusOK).Send(*file.Data)
	}
}

func (h *Handlers) DeleteBinFile() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.BinService.DeleteBin(c); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(Err{
				Error: err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(Message{
			Msg: "файл успешно удален",
		})
	}
}

func (h *Handlers) EditBinName() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.BinService.EditBinName(c); err != nil {
			if strings.Contains(err.Error(), "название файла не может быть") {
				return c.Status(http.StatusBadRequest).JSON(Err{
					Error: err.Error(),
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(Err{
				Error: err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(Message{
			Msg: "имя файла успешно изменено",
		})
	}
}

func selectContentType(contentType string) string {
	prepareHeader := "attachment; filename=file.%s"
	switch {
	case contentType == "image/png":
		return fmt.Sprintf(prepareHeader, "png")
	case contentType == "text/plain; charset=utf-8":
		return fmt.Sprintf(prepareHeader, "txt")
	case contentType == "image/jpeg":
		return fmt.Sprintf(prepareHeader, "jpeg")
	}
	return ""
}
