package ctlv1

import (
	ctlarticles "github.com/diazharizky/go-rest-bootstrap/internal/controllers/v1/articles"
	authctl "github.com/diazharizky/go-rest-bootstrap/internal/controllers/v1/auth"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() (router *fiber.App) {
	authRouter := authctl.NewRouter()
	articlesRouter := ctlarticles.NewRouter()

	router = fiber.New()
	router.Mount("/auth", authRouter)
	router.Mount("/articles", articlesRouter)

	return
}
