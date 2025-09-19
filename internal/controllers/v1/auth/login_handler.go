package authctl

import (
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) LoginHandler(fcx *fiber.Ctx) error {
	statusCode, resp := apiresp.Ok(nil)
	return fcx.Status(statusCode).JSON(resp)
}
