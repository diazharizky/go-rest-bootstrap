package ctlarticles

import (
	"strconv"

	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) List(fcx *fiber.Ctx) error {
	userID := fcx.Query("user_id")
	userIDInt, _ := strconv.Atoi(userID)

	users, err := ctl.ArticleRepository.List(int64(userIDInt))
	if err != nil {
		statusCode, resp := apiresp.CommonError(err)
		return fcx.Status(statusCode).JSON(resp)
	}

	statusCode, resp := apiresp.Ok(users)
	return fcx.Status(statusCode).JSON(resp)
}
