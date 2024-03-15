package interfaces

import "github.com/diazharizky/go-rest-bootstrap/internal/models"

type UserRepository interface {
	List() ([]models.User, error)
	Get(userID int32) (*models.User, error)
	Create(newUser *models.User) error
}

type ArticleRepository interface {
	List(userID int32) ([]models.Article, error)
	Get(articleID int32) (*models.Article, error)
	Create(newArticle *models.Article) error
}

type CreateUserModule interface {
	Call(newUser *models.User) error
}

type CreateArticleModule interface {
	Call(newArticle *models.Article) error
}
