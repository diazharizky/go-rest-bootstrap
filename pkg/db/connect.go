package db

import (
	"database/sql"
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/config"
	_ "github.com/lib/pq"
)

func init() {
	config.Global.SetDefault("db.host", "localhost")
	config.Global.SetDefault("db.port", 5432)
	config.Global.SetDefault("db.user", "rest")
	config.Global.SetDefault("db.password", "rest")
	config.Global.SetDefault("db.db_name", "rest")
}

func GetConnection() *sql.DB {
	host := config.Global.GetString("db.host")
	port := config.Global.GetInt("db.port")
	user := config.Global.GetString("db.user")
	password := config.Global.GetString("db.password")
	dbName := config.Global.GetString("db.db_name")

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
