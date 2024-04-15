package ctlusers

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/interfaces"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
)

type Controller struct {
	UserRepository interfaces.UserRepository
}

func NewDefault() (ctl Controller) {
	ctl.UserRepository = repositories.NewUserRepository()
	return
}
