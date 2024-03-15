package routing

import (
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/config"
	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	controllerv1 "github.com/diazharizky/go-rest-bootstrap/internal/routing/controllers/v1"
	controllerv2 "github.com/diazharizky/go-rest-bootstrap/internal/routing/controllers/v2"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	server *fiber.App
}

func NewRouter() (r router) {
	svr := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	svr.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"ok": true,
		})
	})

	app := core.New()
	api := svr.Group("/api")

	api.Mount("/v1", controllerv1.New(app))
	api.Mount("/v2", controllerv2.New(app))

	r.server = svr

	return
}

func (r router) Start() {
	addr := fmt.Sprintf("%s:%d",
		config.Global.GetString("app.host"),
		config.Global.GetInt("app.port"),
	)

	r.server.Listen(addr)
}
