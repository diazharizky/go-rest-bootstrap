package ctlusers

import (
	"net/http"

	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) Create(fcx *fiber.Ctx) error {
	newUser := new(models.User)
	if err := fcx.BodyParser(newUser); err != nil {
		resp := apiresp.CommonError(err)
		return fcx.Status(http.StatusInternalServerError).JSON(resp)
	}

	if err := ctl.UserRepository.Create(newUser); err != nil {
		resp := apiresp.CommonError(err)
		return fcx.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := apiresp.Success(newUser)
	return fcx.Status(http.StatusCreated).JSON(resp)
}
