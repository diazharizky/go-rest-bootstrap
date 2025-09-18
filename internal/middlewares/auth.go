package middlewares

import (
	"github.com/diazharizky/go-rest-bootstrap/config"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func init() {
	config.Global.SetDefault("db.host", "localhost")
	config.Global.SetDefault("db.port", 5432)
	config.Global.SetDefault("db.user", "gorestbs")
	config.Global.SetDefault("db.password", "gorestbs")
	config.Global.SetDefault("db.name", "gorestbs")
}

func JWTProtected(fcx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(
			config.Global.GetString("APP_JWT_SECRET"),
		)},
		ContextKey: "jwt",
		ErrorHandler: func(fcx *fiber.Ctx, _ error) error {
			statusCode, resp := apiresp.NotAuthenticatedError()
			return fcx.Status(statusCode).JSON(resp)
		},
	})(fcx)
}
