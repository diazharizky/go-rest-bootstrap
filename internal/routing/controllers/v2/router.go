package controllerv2

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/gofiber/fiber/v2"
)

func New(app core.Ctx) (router *fiber.App) {
	router = fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"ok":   true,
			"data": "A warm welcome from `/v2`",
		})
	})

	return
}
