package authctl

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	controller := NewDefault()

	router = fiber.New()
	router.Post("/login", controller.LoginHandler)
	router.Post("/register", controller.RegisterHandler)

	return
}
