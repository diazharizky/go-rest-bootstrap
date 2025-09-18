package ctlarticles

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	controller := NewDefault()

	router = fiber.New()
	router.Get("/", controller.List)
	router.Post("/", middlewares.JWTProtected, controller.Create)

	return
}
