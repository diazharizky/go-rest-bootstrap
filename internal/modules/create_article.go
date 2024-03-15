package modules

import (
	"errors"

	"github.com/diazharizky/go-rest-bootstrap/internal/interfaces"
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
)

type createArticleModule struct {
	userRepository    interfaces.UserRepository
	articleRepository interfaces.ArticleRepository
}

func NewCreateArticleModule() (m createArticleModule) {
	m.userRepository = repositories.NewUserRepository()
	m.articleRepository = repositories.NewArticleRepository()
	return
}

func (m createArticleModule) Call(newArticle *models.Article) error {
	user, err := m.userRepository.Get(newArticle.AuthorID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user is not found")
	}

	if err := m.articleRepository.Create(newArticle); err != nil {
		return err
	}

	return nil
}
