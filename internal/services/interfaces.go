package services

import "github.com/diazharizky/go-rest-bootstrap/internal/models"

type CreateUserService interface {
	Call(newUser *models.User) error
}

type CreateArticleService interface {
	Call(newArticle *models.Article) error
}
