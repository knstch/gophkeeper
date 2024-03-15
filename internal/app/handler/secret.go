package handler

import (
	"fmt"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) StorePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.StoreSecret(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(400).JSON(&Err{
					Error: "поле не может быть пустым",
				})
			}
			fmt.Println(err)
			return c.Status(500).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(202).JSON(&Message{
			Msg: "данны успешно сохранены!",
		})
	}
}

func (h *Handlers) GetAllPrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.SecretService.GetAllSecrets(c)
		if err != nil {
			return c.Status(500).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(200).JSON(data)
	}
}

func (h *Handlers) GetServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.SecretService.GetServiceSecrets(c)
		if err != nil {
			return c.Status(500).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(200).JSON(data)
	}
}

func (h *Handlers) EditServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.EditServiceSecrets(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(400).JSON(&Err{
					Error: "поле не может быть пустым",
				})
			}
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
			Msg: "данные успешно изменены",
		})
	}
}

func (h *Handlers) DeleteServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.DeleteSecrets(c); err != nil {
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
