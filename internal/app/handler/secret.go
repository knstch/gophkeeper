package handler

import (
	"net/http"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) StorePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.StoreSecret(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: "поле не может быть пустым",
				})
			}
			if strings.Contains(err.Error(), common.ErrLength) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(http.StatusCreated).JSON(&Message{
			Msg: "данны успешно сохранены!",
		})
	}
}

func (h *Handlers) GetAllPrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.SecretService.GetAllSecrets(c)
		if err != nil {
			if err == common.ErroNoDataWereFound {
				return c.Status(http.StatusNoContent).JSON(&Err{
					Error: "пусто",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(http.StatusOK).JSON(data)
	}
}

func (h *Handlers) GetServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := h.SecretService.GetSecretsByService(c)
		if err != nil {
			if err == common.ErroNoDataWereFound {
				return c.Status(http.StatusNoContent).JSON(&Err{
					Error: "пусто",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(http.StatusOK).JSON(data)
	}
}

func (h *Handlers) EditServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.EditServiceSecrets(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: "поле не может быть пустым",
				})
			}
			if strings.Contains(err.Error(), common.ErrLength) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			if err == common.ErroNoDataWereFound {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: "ошибка запроса, данные не найдены",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(http.StatusAccepted).JSON(&Message{
			Msg: "данные успешно изменены",
		})
	}
}

func (h *Handlers) DeleteServicePrivates() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.SecretService.DeleteSecrets(c); err != nil {
			if err == common.ErroNoDataWereFound {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: "ошибка запроса, данные не найдены",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: "внутренняя ошибка сервиса",
			})
		}
		return c.Status(http.StatusOK).JSON(&Message{
			Msg: "данные успешно удалены",
		})
	}
}
