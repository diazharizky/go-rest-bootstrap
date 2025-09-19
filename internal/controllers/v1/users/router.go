package ctlusers

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	controller := NewDefault()

	router = fiber.New()
	router.Get("/", controller.ListHandler)
	router.Post("/", controller.CreateHandler)

	return
}
