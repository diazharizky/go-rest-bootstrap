package db

import (
	"fmt"

	"gorm.io/gorm"
)

func GetConnection() (conn *gorm.DB, err error) {
	db := NewPostgres()

	conn, err = db.Connect()
	if err != nil {
		return nil, fmt.Errorf("error unable to get database connection: %v", err)
	}

	return
}
