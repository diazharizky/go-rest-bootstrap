package ctlusers

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) Create(fcx *fiber.Ctx) error {
	newUser := new(models.User)
	if err := fcx.BodyParser(newUser); err != nil {
		statusCode, resp := apiresp.CommonError(err)
		return fcx.Status(statusCode).JSON(resp)
	}

	if err := ctl.UserRepository.Create(newUser); err != nil {
		statusCode, resp := apiresp.CommonError(err)
		return fcx.Status(statusCode).JSON(resp)
	}

	statusCode, resp := apiresp.Ok(newUser)
	return fcx.Status(statusCode).JSON(resp)
}
