package handler

import (
	"net/http"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) RegisterWithEmail() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken, err := h.AuthService.SignUp(c.Context(), c.Request().Body())
		if err != nil {
			if err.Error() == common.ErrIntegrityViolation.Error() {
				return c.Status(http.StatusConflict).JSON(&Err{
					Error: "эта почта уже занята",
				})
			}
			if strings.Contains(err.Error(), common.ErrBadPass.Error()) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			if strings.Contains(err.Error(), common.ErrBadEmail.Error()) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			if strings.Contains(err.Error(), common.ErrFieldIsEmpty.Error()) {
				return c.Status(http.StatusBadRequest).JSON(&Err{
					Error: err.Error(),
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(&Err{
				Error: err.Error(),
			})

		}

		c.Cookie(&fiber.Cookie{
			Name:  "auth",
			Value: accessToken,
			Path:  "/",
		})

		return c.Status(http.StatusOK).JSON(&Message{
			Msg: "вы успешно зарегестрировались!",
		})
	}
}

func (h *Handlers) AuthenticateWithEmail() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken, err := h.AuthService.SignIn(c.Request().Body())
		if err != nil {
			if err.Error() == common.ErrUserNotFound.Error() {
				return c.Status(http.StatusNotFound).JSON(&Message{
					Msg: "неверная почта или пароль",
				})
			}
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:  "auth",
			Value: accessToken,
			Path:  "/",
		})

		return c.Status(http.StatusOK).JSON(&Message{
			Msg: "вы успешно залогинились!",
		})
	}
}
