package server

import (
	"errors"
	"fmt"
	"net/http"
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
	engine *fiber.App
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
		Max:        config.Global.GetInt("app.throttling.max.requests"),
		Expiration: config.Global.GetDuration("app.throttling.expiration") * time.Second,
		LimitReached: func(fcx *fiber.Ctx) error {
			resp := apiresp.CommonError(
				errors.New("too many requests"),
			)
			return fcx.Status(http.StatusTooManyRequests).JSON(resp)
		},
	}))

	app.Get("/healthcheck", func(fcx *fiber.Ctx) error {
		return fcx.
			Status(http.StatusOK).
			JSON(apiresp.Success(nil))
	})

	apiPath := app.Group("/api")
	apiPath.Mount("/v1", ctlv1.NewRouter())
	apiPath.Mount("/v2", ctlv2.NewRouter())

	svr.engine = app

	return
}

func (svr server) Start() {
	addr := fmt.Sprintf("%s:%d",
		config.Global.GetString("app.host"),
		config.Global.GetInt("app.port"),
	)

	svr.engine.Listen(addr)
}
