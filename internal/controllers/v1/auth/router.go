package authctl

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	controller := NewDefault()

	router = fiber.New()
	router.Get("/login", controller.Login)
	router.Post("/register", controller.Register)

	return
}
