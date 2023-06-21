package users

import (
	"net/http"

	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl controller) List(ctx *fiber.Ctx) error {
	users, err := ctl.core.UserRepository.List()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	return ctx.JSON(users)
}
