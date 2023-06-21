package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func NewCustomMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Next()
	}
}
