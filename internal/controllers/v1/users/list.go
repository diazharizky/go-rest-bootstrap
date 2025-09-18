package ctlusers

import (
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) List(fcx *fiber.Ctx) error {
	users, err := ctl.UserRepository.List()
	if err != nil {
		statusCode, resp := apiresp.UnknownError(err)
		return fcx.Status(statusCode).JSON(resp)
	}

	statusCode, resp := apiresp.Ok(users)
	return fcx.Status(statusCode).JSON(resp)
}
