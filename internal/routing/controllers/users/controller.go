package users

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	core *core.Core
}

func RegisterController(router fiber.Router, xore *core.Core) {
	routes := router.Group("/users")

	ctl := controller{
		core: xore,
	}

	routes.Get("/", ctl.List)
	routes.Get("/:username", ctl.Get)
	routes.Post("/", ctl.Create)
}
