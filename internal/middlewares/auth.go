package middlewares

import (
	"github.com/diazharizky/go-rest-bootstrap/config"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func JWTProtected(fcx *fiber.Ctx) error {
	appSecret := config.Global.GetString("app.jwt_secret")
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(appSecret)},
		ContextKey: "jwt",
		ErrorHandler: func(fcx *fiber.Ctx, _ error) error {
			statusCode, resp := apiresp.NotAuthenticatedError()
			return fcx.Status(statusCode).JSON(resp)
		},
	})(fcx)
}
