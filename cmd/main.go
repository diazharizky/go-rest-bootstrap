package main

import (
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/internal/server"
	"github.com/diazharizky/go-rest-bootstrap/pkg/db"
)

func main() {
	conn := db.MustGetConnection()

	if err := conn.AutoMigrate(&models.User{}, &models.Article{}); err != nil {
		panic(
			fmt.Sprintf("failed to perform migration: %v", err),
		)
	}

	svr := server.New()
	svr.Start()
}
