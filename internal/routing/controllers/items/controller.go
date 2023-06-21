package items

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	core *core.Core
}

func RegisterController(router fiber.Router, xore *core.Core) {
	routes := router.Group("/items")

	ctl := controller{
		core: xore,
	}

	routes.Get("/", ctl.Get)
}
