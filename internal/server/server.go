package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/diazharizky/go-rest-bootstrap/config"
	ctlv1 "github.com/diazharizky/go-rest-bootstrap/internal/controllers/v1"
	ctlv2 "github.com/diazharizky/go-rest-bootstrap/internal/controllers/v2"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type server struct {
	app *fiber.App
}

func init() {
	config.Global.SetDefault("app.host", "localhost")
	config.Global.SetDefault("app.port", 8080)
	config.Global.SetDefault("app.throttling.max.requests", 20)
	config.Global.SetDefault("app.throttling.expiration", 30)
}

func New() (svr server) {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		// This rate limiter might not be working
		// when the app is running in multiple nodes.
		// Probably we need to utilise in-memory database
		// to store the metrics.

		Max:        config.Global.GetInt("app.throttling.max.requests"),
		Expiration: config.Global.GetDuration("app.throttling.expiration") * time.Second,
		LimitReached: func(fcx *fiber.Ctx) error {
			statusCode, resp := apiresp.CommonError(errors.New("too many requests"))
			return fcx.
				Status(statusCode).
				JSON(resp)
		},
	}))

	app.Get("/healthcheck", func(fcx *fiber.Ctx) error {
		statusCode, resp := apiresp.Ok(nil)
		return fcx.Status(statusCode).JSON(resp)
	})

	apiBasePath := app.Group("/api")
	apiBasePath.Mount("/v1", ctlv1.NewRouter())
	apiBasePath.Mount("/v2", ctlv2.NewRouter())

	svr.app = app

	return
}

func (svr server) Start() {
	addr := fmt.Sprintf("%s:%d",
		config.Global.GetString("app.host"),
		config.Global.GetInt("app.port"),
	)

	svr.app.Listen(addr)
}
