package approuter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/internal/app/common"
	"github.com/knstch/gophkeeper/internal/app/handler"
	"github.com/knstch/gophkeeper/internal/middleware"
)

func InitRouter(app *fiber.App, handlers *handler.Handlers, storage common.Storager) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.RegisterWithEmail())
	auth.Post("/", handlers.AuthenticateWithEmail())

	secret := app.Group("/secret", middleware.WithCookieLogin())
	secret.Post("/", handlers.StorePrivates())
	secret.Get("/", handlers.GetAllPrivates())
	secret.Get("/:service", handlers.GetServicePrivates())
	secret.Put("/", handlers.EditServicePrivates())
}
