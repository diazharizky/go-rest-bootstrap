package ctlv1

import (
	ctlarticles "github.com/diazharizky/go-rest-bootstrap/internal/controllers/v1/articles"
	ctlusers "github.com/diazharizky/go-rest-bootstrap/internal/controllers/v1/users"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	usersRouter := ctlusers.NewRouter()
	articlesRouter := ctlarticles.NewRouter()

	router = fiber.New()
	router.Mount("/users", usersRouter)
	router.Mount("/articles", articlesRouter)

	return
}
