package ctlarticles

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	controller := NewDefault()

	router = fiber.New()
	router.Get("/", controller.ListHandler)
	router.Post("/", middlewares.JWTProtected, controller.CreateHandler)

	return
}
