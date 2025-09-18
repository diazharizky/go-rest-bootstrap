package ctlarticles

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) Create(fcx *fiber.Ctx) error {
	newArticle := new(models.Article)
	if err := fcx.BodyParser(newArticle); err != nil {
		statusCode, resp := apiresp.UnknownError(err)
		return fcx.Status(statusCode).JSON(resp)
	}

	if err := ctl.CreateArticleService.Call(newArticle); err != nil {
		statusCode, resp := apiresp.UnknownError(err)
		return fcx.Status(statusCode).JSON(resp)
	}

	statusCode, resp := apiresp.Ok(newArticle)
	return fcx.Status(statusCode).JSON(resp)
}
