package authctl

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
	"github.com/diazharizky/go-rest-bootstrap/internal/services"
)

type Controller struct {
	UserRepository  repositories.UserRepository
	RegisterService services.RegisterService
}

func NewDefault() (ctl Controller) {
	ctl.UserRepository = repositories.NewUserRepository()
	ctl.RegisterService = services.NewRegisterService()
	return
}
