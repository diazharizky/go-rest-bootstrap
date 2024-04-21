package repositories

import (
	"errors"

	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/db"
	"gorm.io/gorm"
)

type userRepository struct {
	tableName string
}

func NewUserRepository() (r userRepository) {
	r.tableName = "users"
	return
}

func (r userRepository) List() ([]models.User, error) {
	conn := db.MustGetConnection()

	users := []models.User{}
	if tx := conn.Find(&users); tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r userRepository) Get(userID int64) (*models.User, error) {
	conn := db.MustGetConnection()

	user := models.User{}
	if tx := conn.First(&user, userID); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepository) Create(newUser *models.User) error {
	conn := db.MustGetConnection()

	if tx := conn.Create(&newUser); tx.Error != nil {
		return tx.Error
	}

	return nil
}
