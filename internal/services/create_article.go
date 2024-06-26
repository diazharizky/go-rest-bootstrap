package services

import (
	"errors"

	"github.com/diazharizky/go-rest-bootstrap/internal/interfaces"
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
)

type createArticleService struct {
	UserRepository    interfaces.UserRepository
	ArticleRepository interfaces.ArticleRepository
}

func NewCreateArticleService() (m createArticleService) {
	m.UserRepository = repositories.NewUserRepository()
	m.ArticleRepository = repositories.NewArticleRepository()
	return
}

func (m createArticleService) Call(newArticle *models.Article) error {
	user, err := m.UserRepository.Get(newArticle.AuthorID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user is not found")
	}

	if err := m.ArticleRepository.Create(newArticle); err != nil {
		return err
	}

	return nil
}
