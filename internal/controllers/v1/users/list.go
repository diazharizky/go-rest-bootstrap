package ctlusers

import (
	"net/http"

	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) List(fcx *fiber.Ctx) error {
	users, err := ctl.UserRepository.List()
	if err != nil {
		resp := apiresp.CommonError(err)
		return fcx.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := apiresp.Success(users)
	return fcx.Status(http.StatusOK).JSON(resp)
}
