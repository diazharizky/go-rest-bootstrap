package ctlarticles

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/interfaces"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
	"github.com/diazharizky/go-rest-bootstrap/internal/services"
)

type Controller struct {
	ArticleRepository    interfaces.ArticleRepository
	CreateArticleService interfaces.CreateArticleService
}

func NewDefault() (ctl Controller) {
	ctl.ArticleRepository = repositories.NewArticleRepository()
	ctl.CreateArticleService = services.NewCreateArticleService()
	return
}
