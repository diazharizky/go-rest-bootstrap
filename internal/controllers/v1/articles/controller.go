package ctlarticles

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
	"github.com/diazharizky/go-rest-bootstrap/internal/services"
)

type Controller struct {
	ArticleRepository    repositories.ArticleRepository
	CreateArticleService services.CreateArticleService
}

func NewDefault() (ctl Controller) {
	ctl.ArticleRepository = repositories.NewArticleRepository()
	ctl.CreateArticleService = services.NewCreateArticleService()
	return
}
