package db

import (
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustGetConnection() *gorm.DB {
	host := config.Global.GetString("db.host")
	port := config.Global.GetInt("db.port")
	user := config.Global.GetString("db.user")
	password := config.Global.GetString("db.password")
	dbName := config.Global.GetString("db.name")

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(
			fmt.Sprintf("failed to connect database: %v", err),
		)
	}

	return db
}
