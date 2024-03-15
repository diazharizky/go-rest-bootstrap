package controllerv1

import (
	"net/http"
	"strconv"

	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

type articlesController struct {
	app core.Ctx
}

func NewArticlesController(app core.Ctx) (ctl articlesController) {
	ctl.app = app
	return
}

func (ctl articlesController) List(c *fiber.Ctx) error {
	userID := c.Params(routeParamUserID)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	articles, err := ctl.app.ArticleRepository.List(int32(userIDInt))
	if err != nil {
		return err
	}

	return c.
		Status(http.StatusOK).
		JSON(
			apiresp.Success(articles),
		)
}

func (ctl articlesController) Get(c *fiber.Ctx) error {
	return nil
}

func (ctl articlesController) Create(c *fiber.Ctx) error {
	newArticle := new(models.Article)
	if err := c.BodyParser(newArticle); err != nil {
		return err
	}

	userID := c.Params(routeParamUserID)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	newArticle.AuthorID = int32(userIDInt)

	if err := ctl.app.CreateArticleModule.Call(newArticle); err != nil {
		return err
	}

	return c.
		Status(http.StatusCreated).
		JSON(
			apiresp.Success(newArticle),
		)
}
