package authctl

import (
	"errors"
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/diazharizky/go-rest-bootstrap/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type loginReqBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (ctl Controller) LoginHandler(fcx *fiber.Ctx) error {
	var reqBody loginReqBody
	if err := fcx.BodyParser(&reqBody); err != nil {
		statusCode, resp := apiresp.FatalError()
		return fcx.Status(statusCode).JSON(resp)
	}

	validateInput := utils.NewInputValidator()
	if err := validateInput.Struct(reqBody); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			xValidationErrors := utils.XValidationErrors(validationErrors)
			statusCode, resp := apiresp.InputRequiredError(xValidationErrors.ResolveInputErrors())
			return fcx.Status(statusCode).JSON(resp)
		}

		statusCode, resp := apiresp.UnknownError(err)
		return fcx.Status(statusCode).JSON(resp)
	}

	filter := map[string]any{"email": reqBody.Email}
	user, err := ctl.UserRepository.GetBy(filter)
	if err != nil {
		statusCode, resp := apiresp.FatalError()
		return fcx.Status(statusCode).JSON(resp)
	}

	if user == nil {
		statusCode, resp := apiresp.NotFoundError()
		return fcx.Status(statusCode).JSON(resp)
	}

	isMatched := utils.ComparePassword(*user.Password, reqBody.Password)
	if !isMatched {
		statusCode, resp := apiresp.CredentialsError()
		return fcx.Status(statusCode).JSON(resp)
	}

	token, err := utils.GenerateToken(uint(user.ID))
	if err != nil {
		fmt.Println("Error:", err.Error())
		statusCode, resp := apiresp.FatalError()
		return fcx.Status(statusCode).JSON(resp)
	}

	statusCode, resp := apiresp.Ok(map[string]any{"token": token})
	return fcx.Status(statusCode).JSON(resp)
}
