package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/knstch/gophkeeper/cmd/config"
	"github.com/knstch/gophkeeper/internal/app/approuter"
	"github.com/knstch/gophkeeper/internal/app/handler"
	"github.com/knstch/gophkeeper/internal/app/storage/psql"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config.ParseConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	psqlStorage, err := psql.NewPsqlStorage(config.ReadyConfig.DSN)
	if err != nil {
		return err
	}

	handlers := handler.NewHandler(psqlStorage)

	approuter.InitRouter(app, handlers, psqlStorage)

	return app.Listen(config.ReadyConfig.ServerAddr)
}
