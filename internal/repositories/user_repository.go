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

func NewUserRepository() (rep userRepository) {
	rep.tableName = "users"
	return
}

func (rep userRepository) List() ([]models.User, error) {
	con := db.MustGetConnection()

	users := []models.User{}
	if trx := con.Find(&users); trx.Error != nil {
		return nil, trx.Error
	}

	return users, nil
}

func (rep userRepository) Get(userID int64) (*models.User, error) {
	con := db.MustGetConnection()

	var user models.User
	if trx := con.First(&user, userID); trx.Error != nil {
		if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, trx.Error
	}

	return &user, nil
}

func (rep userRepository) GetBy(filter map[string]any) (*models.User, error) {
	con := db.MustGetConnection()

	var user models.User
	if trx := con.Where(filter).First(user); trx.Error != nil {
		if errors.Is(trx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, trx.Error
	}

	return &user, nil
}

func (rep userRepository) Create(newUser *models.User) error {
	con := db.MustGetConnection()

	trx := con.Create(&newUser)
	if trx.Error != nil {
		return trx.Error
	}
	newUser.ID = trx.RowsAffected

	return nil
}
