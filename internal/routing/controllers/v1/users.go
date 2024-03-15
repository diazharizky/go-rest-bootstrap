package controllerv1

import (
	"net/http"
	"strconv"

	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	app core.Ctx
}

func NewUsersController(app core.Ctx) (ctl usersController) {
	ctl.app = app
	return
}

func (ctl usersController) List(c *fiber.Ctx) error {
	users, err := ctl.app.UserRepository.List()
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(
			apiresp.Success(users),
		)
}

func (ctl usersController) Get(c *fiber.Ctx) error {
	userID := c.Params(routeParamUserID)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	user, err := ctl.app.UserRepository.Get(int32(userIDInt))
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(
			apiresp.Success(user),
		)
}

func (ctl usersController) Create(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return err
	}

	if err := ctl.app.CreateUserModule.Call(newUser); err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(
			apiresp.Success(newUser),
		)
}
