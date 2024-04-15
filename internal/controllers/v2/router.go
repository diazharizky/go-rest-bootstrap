package ctlv2

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	router = fiber.New()
	return
}
