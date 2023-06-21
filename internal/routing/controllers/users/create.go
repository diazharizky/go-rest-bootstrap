package users

import (
	"net/http"

	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl controller) Create(ctx *fiber.Ctx) error {
	newUser := &models.User{}

	if err := ctx.BodyParser(newUser); err != nil {
		return err
	}

	if err := ctl.core.UserRepository.Create(newUser); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	return ctx.JSON(newUser)
}
