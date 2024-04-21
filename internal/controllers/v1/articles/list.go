package ctlarticles

import (
	"net/http"
	"strconv"

	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) List(fcx *fiber.Ctx) error {
	userID := fcx.Query("user_id")
	userIDInt, _ := strconv.Atoi(userID)

	users, err := ctl.ArticleRepository.List(int64(userIDInt))
	if err != nil {
		resp := apiresp.CommonError(err)
		return fcx.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := apiresp.Success(users)
	return fcx.Status(http.StatusOK).JSON(resp)
}
