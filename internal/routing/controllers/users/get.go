package users

import (
	"net/http"

	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl controller) Get(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	user, err := ctl.core.UserRepository.Get(username)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	return ctx.JSON(user)
}
