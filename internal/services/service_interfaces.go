package services

import "github.com/diazharizky/go-rest-bootstrap/internal/models"

type CreateUserService interface {
	Execute(newUser *models.User) error
}

type CreateArticleService interface {
	Execute(newArticle *models.Article) error
}

type RegisterService interface {
	Execute(newUser *models.User) error
}
