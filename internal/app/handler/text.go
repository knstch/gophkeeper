package handler

import (
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) StoreText() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.TextService.StoreTexts(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(400).JSON(&Err{
					Error: err.Error(),
				})
			}
			if strings.Contains(err.Error(), common.ErrLength) {
				return c.Status(400).JSON(&Err{
					Error: err.Error(),
				})
			}
			return c.Status(500).JSON(&Err{
				Error: err.Error(),
			})
		}
		return nil
	}
}

func (h *Handlers) GetAllTexts() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.TextService.GetAllTexts(c)
		if err != nil {
			return c.Status(500).JSON(&Err{
				Error: err.Error(),
			})
		}
		if len(data.Texts) == 0 {
			return c.Status(204).JSON(&Err{
				Error: "пусто",
			})
		}
		return c.Status(200).JSON(data)
	}
}

func (h *Handlers) GetTextByTitle() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.TextService.GetTextByTitle(c)
		if err != nil {
			if err == common.ErroNoDataWereFound {
				return c.Status(204).JSON(data)
			}
			return c.Status(500).JSON(&Err{
				Error: err.Error(),
			})
		}
		return c.Status(200).JSON(data)
	}
}

func (h *Handlers) EditText() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.TextService.EditText(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(400).JSON(&Err{
					Error: "поле не может быть пустым",
				})
			}
			if strings.Contains(err.Error(), common.ErrLength) {
				return c.Status(400).JSON(&Err{
					Error: err.Error(),
				})
			}
			if err == common.ErroNoDataWereFound {
				return c.Status(400).JSON(&Err{
					Error: "ошибка запроса, данные не найдены",
				})
			}
			return c.Status(500).JSON(&Err{
				Error: err.Error(),
			})
		}
		return c.Status(200).JSON(&Message{
			Msg: "текст изменен",
		})
	}
}

func (h *Handlers) DeleteText() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.TextService.DeleteText(c); err != nil {
			if err == common.ErroNoDataWereFound {
				return c.Status(400).JSON(&Err{
					Error: "ошибка запроса, данные не найдены",
				})
			}
			return c.Status(500).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(200).JSON(&Message{
			Msg: "данные успешно удалены",
		})
	}
}
