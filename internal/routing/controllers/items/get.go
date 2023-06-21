package items

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (controller) Get(ctx *fiber.Ctx) error {
	item := models.Item{
		ID:   1,
		Name: "Fender telecaster",
	}

	return ctx.JSON(item)
}
