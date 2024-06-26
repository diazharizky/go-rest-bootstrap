package interfaces

import "github.com/diazharizky/go-rest-bootstrap/internal/models"

type UserRepository interface {
	List() ([]models.User, error)
	Get(userID int64) (*models.User, error)
	Create(newUser *models.User) error
}

type ArticleRepository interface {
	List(userID int64) ([]models.Article, error)
	Get(articleID int64) (*models.Article, error)
	Create(newArticle *models.Article) error
}

type CreateUserService interface {
	Call(newUser *models.User) error
}

type CreateArticleService interface {
	Call(newArticle *models.Article) error
}
