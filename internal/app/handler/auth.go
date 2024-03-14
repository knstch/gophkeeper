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
		accessToken, err := h.AuthService.SignUp(c.Params("email"), c.Params("password"))
		if err.Error() == common.ErrIntegrityViolation {
			return c.Status(409).JSON(&Message{
				Msg: "This email is already taken",
			})
		}
		if err != nil {
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
		accessToken, err := h.AuthService.SignIn(c.Params("email"), c.Params("password"))
		if err.Error() == common.ErrUserNotFound {
			return c.Status(404).JSON(&Message{
				Msg: "This user is not found",
			})
		}
		if err != nil {
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
