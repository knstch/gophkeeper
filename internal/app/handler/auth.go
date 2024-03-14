package handler

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
)

type Message struct {
	Msg string `json:"message"`
}

func (h *Handlers) RegisterWithEmail() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken, err := h.AuthService.SignUp(c.Context(), c.Request().Body())
		if err != nil {
			if err.Error() == common.ErrIntegrityViolation {
				return c.Status(409).JSON(&Message{
					Msg: "This email is already taken",
				})
			}

			if err != nil {
				return c.Status(400).JSON(&Message{
					Msg: err.Error(),
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
			Msg: "You have successfully signed up!",
		})
	}
}

func (h *Handlers) AuthenticateWithEmail() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken, err := h.AuthService.SignIn(c.Request().Body())
		if err != nil {
			if err.Error() == common.ErrUserNotFound {
				return c.Status(404).JSON(&Message{
					Msg: "Wrong email or password",
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
			Msg: "You have successfully signed in!",
		})
	}
}
