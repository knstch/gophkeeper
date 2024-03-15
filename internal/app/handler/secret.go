package handler

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) StorePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.StoreSecret(c, c.Body()); err != nil {
			if err == common.ErrNotLoggedIn {
				return c.Status(401).JSON(Message{
					Msg: "необходимо зарегестрироваться или авторизоваться",
				})
			}
			return c.Status(500).JSON(Message{
				Msg: "внутренняя ошибка сервиса",
			})
		}

		return c.Status(202).JSON(Message{
			Msg: "данны успешно сохранены!",
		})
	}
}

func (h *Handlers) GetAllPrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.SecretService.GetAllSecrets(c)
		if err != nil {
			if err == common.ErrNotLoggedIn {
				return c.Status(401).JSON(Message{
					Msg: "необходимо зарегестрироваться или авторизоваться",
				})
			}
			return c.Status(500).JSON(Message{
				Msg: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(200).JSON(data)
	}
}

func (h *Handlers) GetServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.SecretService.GetServiceSecrets(c)
		if err != nil {
			if err == common.ErrNotLoggedIn {
				return c.Status(401).JSON(Message{
					Msg: "необходимо зарегестрироваться или авторизоваться",
				})
			}
			return c.Status(500).JSON(Message{
				Msg: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(200).JSON(data)
	}
}

func (h *Handlers) EditServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.EditServiceSecrets(c); err != nil {
			if err == common.ErrNotLoggedIn {
				return c.Status(401).JSON(Message{
					Msg: "необходимо зарегестрироваться или авторизоваться",
				})
			}
			if err == common.ErroNoDataWereFound {
				return c.Status(400).JSON(Message{
					Msg: "ошибка запроса, данные не найдены",
				})
			}
			return c.Status(500).JSON(Message{
				Msg: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(200).JSON(Message{
			Msg: "данные успешно изменены",
		})
	}
}
