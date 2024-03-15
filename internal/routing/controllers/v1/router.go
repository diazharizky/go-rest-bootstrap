package controllerv1

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/gofiber/fiber/v2"
)

const (
	routeParamUserID    = "user_id"
	routeParamArticleID = "article_id"
)

func New(app core.Ctx) (router *fiber.App) {
	router = fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	usersController := NewUsersController(app)

	usersPath := router.Group("/users")
	usersPath.Get("/", usersController.List)
	usersPath.Post("/", usersController.Create)

	userPath := usersPath.Group("/:" + routeParamUserID)
	userPath.Get("/", usersController.Get)

	articlesController := NewArticlesController(app)

	userArticlesPath := userPath.Group("/articles")
	userArticlesPath.Get("/", articlesController.List)
	userArticlesPath.Post("/", articlesController.Create)

	userArticlePath := userArticlesPath.Group("/:" + routeParamArticleID)
	userArticlePath.Get("/", articlesController.Get)

	return
}
