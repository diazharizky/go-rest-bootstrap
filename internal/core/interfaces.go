package core

import "github.com/diazharizky/go-rest-bootstrap/internal/models"

type IUserRepository interface {
	List() ([]models.User, error)
	Get(username string) (*models.User, error)
	Create(newUser *models.User) error
}

type IItemRepository interface {
	List() ([]models.Item, error)
}
