package handler

import (
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

func (h *Handlers) RegisterWithEmail() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken, err := h.AuthService.SignUp(c.Context(), c.Request().Body())
		if err != nil {
			if err.Error() == common.ErrIntegrityViolation.Error() {
				return c.Status(409).JSON(&Err{
					Error: "эта почта уже занята",
				})
			}
			if strings.Contains(err.Error(), common.ErrBadPass.Error()) {
				return c.Status(400).JSON(&Err{
					Error: err.Error(),
				})
			}
			if strings.Contains(err.Error(), common.ErrBadEmail.Error()) {
				return c.Status(400).JSON(&Err{
					Error: err.Error(),
				})
			}
			if err != nil {
				return c.Status(400).JSON(&Err{
					Error: err.Error(),
				})
			}

			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:  "auth",
			Value: accessToken,
			Path:  "/",
		})

		return c.Status(200).JSON(&Message{
			Msg: "вы успешно залогинились!",
		})
	}
}

func (h *Handlers) AuthenticateWithEmail() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken, err := h.AuthService.SignIn(c.Request().Body())
		if err != nil {
			if err.Error() == common.ErrUserNotFound.Error() {
				return c.Status(404).JSON(&Message{
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

		return c.Status(200).JSON(&Message{
			Msg: "вы успешно залогинились!",
		})
	}
}
