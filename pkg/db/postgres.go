package db

import (
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgql struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

func init() {
	config.Global.SetDefault("db.host", "0.0.0.0")
	config.Global.SetDefault("db.port", "5432")
	config.Global.SetDefault("db.user", "gofiber")
	config.Global.SetDefault("db.password", "gofiber")
	config.Global.SetDefault("db.database", "gofiber")
	config.Global.SetDefault("db.sslmode", "disable")
}

func NewPostgres() pgql {
	return pgql{
		Host:     config.Global.GetString("db.host"),
		Port:     config.Global.GetString("db.port"),
		User:     config.Global.GetString("db.user"),
		Password: config.Global.GetString("db.password"),
		DbName:   config.Global.GetString("db.database"),
		SslMode:  config.Global.GetString("db.sslmode"),
	}
}

func (pg pgql) Connect() (*gorm.DB, error) {
	dsn := pg.dsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error unable to connect to PostgreSQL DB: %v", err)
	}

	return db, nil
}

func (pg pgql) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg.Host,
		pg.Port,
		pg.User,
		pg.Password,
		pg.DbName,
		pg.SslMode,
	)
}
