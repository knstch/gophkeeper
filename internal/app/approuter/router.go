package approuter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/handler"
)

func InitRouter(app *fiber.App, handlers *handler.Handlers, storage common.Storager) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.RegisterWithEmail())
	auth.Post("/", handlers.AuthenticateWithEmail())
}
