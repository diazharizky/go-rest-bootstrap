package ctlusers

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
)

type Controller struct {
	UserRepository repositories.UserRepository
}

func NewDefault() (ctl Controller) {
	ctl.UserRepository = repositories.NewUserRepository()
	return
}
