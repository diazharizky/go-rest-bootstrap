package main

import (
	"log"

	"github.com/diazharizky/go-rest-bootstrap/internal/core"
	"github.com/diazharizky/go-rest-bootstrap/internal/routing"
)

func main() {
	xore, err := core.New()
	if err != nil {
		log.Fatalf("Encounter error: %v", err)
	}

	router := routing.NewRouter(xore)

	router.Start()
}
