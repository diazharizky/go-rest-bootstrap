package repositories

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		db: db,
	}
}

func (rep userRepository) List() (users []models.User, err error) {
	if err = rep.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return
}

func (rep userRepository) Get(username string) (user *models.User, err error) {
	if err = rep.db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}

	return
}

func (rep userRepository) Create(newUser *models.User) error {
	if err := rep.db.Create(newUser).Error; err != nil {
		return err
	}

	return nil
}
