package routing

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/diazharizky/go-rest-bootstrap/internal/routing/controllers/items"
	"github.com/diazharizky/go-rest-bootstrap/internal/routing/controllers/users"
	"github.com/diazharizky/go-rest-bootstrap/internal/routing/middlewares"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	server *fiber.App
}

func NewRouter(xore *core.Core) (r router) {
	svr := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	api := svr.Group("api")

	v1 := api.Group("v1")
	{
		v1.Use(middlewares.NewCustomMiddleware())

		users.RegisterController(v1, xore)
		items.RegisterController(v1, xore)
	}

	r.server = svr

	return
}

func (r router) Start() {
	r.server.Listen(":3000")
}
