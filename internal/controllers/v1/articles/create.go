package ctlarticles

import (
	"net/http"

	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl Controller) Create(fcx *fiber.Ctx) error {
	newArticle := new(models.Article)
	if err := fcx.BodyParser(newArticle); err != nil {
		resp := apiresp.CommonError(err)
		return fcx.Status(http.StatusInternalServerError).JSON(resp)
	}

	if err := ctl.CreateArticleService.Call(newArticle); err != nil {
		resp := apiresp.CommonError(err)
		return fcx.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := apiresp.Success(newArticle)
	return fcx.Status(http.StatusCreated).JSON(resp)
}
