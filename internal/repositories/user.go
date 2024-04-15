package repositories

import (
	"database/sql"
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/db"
)

type userRepository struct {
	tableName string
}

func NewUserRepository() (r userRepository) {
	r.tableName = "users"
	return
}

func (r userRepository) List() ([]models.User, error) {
	db := db.GetConnection()
	defer db.Close()

	rows, err := db.Query(
		fmt.Sprintf("SELECT id, email, full_name, created_at, deleted_at FROM %s", r.tableName),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}

	user := models.User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.FullName, &user.CreatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r userRepository) Get(userID int64) (*models.User, error) {
	db := db.GetConnection()
	defer db.Close()

	row := db.QueryRow(
		fmt.Sprintf("SELECT id, email, full_name, created_at, deleted_at FROM %s WHERE id=$1", r.tableName),
		userID,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	user := models.User{}

	err := row.Scan(&user.ID, &user.Email, &user.FullName, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r userRepository) Create(newUser *models.User) error {
	db := db.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(
		fmt.Sprintf("INSERT INTO %s(email, full_name) VALUES($1, $2) RETURNING id, created_at", r.tableName),
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newUser.Email, newUser.FullName).Scan(&newUser.ID, &newUser.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
