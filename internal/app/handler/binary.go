package handler

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

func (h *Handlers) StoreBin() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Println(http.DetectContentType(c.Body()))
		c.Set(fiber.HeaderContentDisposition, "attachment; filename=file.png")
		return c.Status(http.StatusOK).Send(c.Body())
	}
}
