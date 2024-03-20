package handler

import (
	"net/http"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) StoreBankCard() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.BankService.StoreCard(c); err != nil {
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
				Error: err.Error(),
			})
		}
		return c.Status(http.StatusCreated).JSON(Message{
			Msg: "данные сохранены!",
		})
	}
}

func (h *Handlers) EditBankCard() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.BankService.EditBankCard(c); err != nil {
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			if strings.Contains(err.Error(), common.ErrLength) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			if err == common.ErrNoDataWereFound {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: "ошибка запроса, данные не найдены",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: err.Error(),
			})
		}

		return c.Status(http.StatusAccepted).JSON(Message{
			Msg: "данные изменены",
		})
	}
}

func (h *Handlers) DeleteBankCard() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if err := h.BankService.DeleteBankCard(c); err != nil {
			if err == common.ErrNoDataWereFound {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: "ошибка запроса, данные не найдены",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(Message{
			Msg: "данные удалены",
		})
	}
}

func (h *Handlers) GetAllCards() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		cards, err := h.BankService.GetAllCards(c)
		if err != nil {
			if err == common.ErrNoDataWereFound {
				return c.Status(http.StatusNoContent).JSON(&Err{
					Error: "пусто",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(cards)
	}
}

func (h *Handlers) GetCardsByBankName() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		cards, err := h.BankService.GetCardsByBankName(c)
		if err != nil {
			if err == common.ErrNoDataWereFound {
				return c.Status(http.StatusNoContent).JSON(&Err{
					Error: "пусто",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(cards)
	}
}
