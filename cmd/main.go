package main

import (
	"github.com/diazharizky/go-rest-bootstrap/config"
	"github.com/diazharizky/go-rest-bootstrap/internal/routing"
)

func init() {
	config.Global.SetDefault("mongodb.db", "gorestbootstrap")
}

func main() {
	router := routing.NewRouter()
	router.Start()
}
