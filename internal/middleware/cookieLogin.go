package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/handler"
	"github.com/knstch/gophkeeper/internal/app/service/jwt"
)

func WithCookieLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req, err := adaptor.ConvertRequest(c, false)
		if err != nil {
			return err
		}

		userEmail, err := jwt.GetEmail(req)
		if err != nil {
			if err == common.ErrNotLoggedIn || err == common.ErrInvalidToken {
				return c.Status(http.StatusForbidden).JSON(handler.Message{
					Msg: "необходимо зарегестрироваться или авторизоваться",
				})
			}
			return err
		}

		c.Locals("login", userEmail)

		return c.Next()
	}
}
