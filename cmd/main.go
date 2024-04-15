package main

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/server"
)

func main() {
	svr := server.New()
	svr.Start()
}
