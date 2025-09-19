package authctl

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/services"
)

type Controller struct {
	RegisterService services.RegisterService
}

func NewDefault() (ctl Controller) {
	ctl.RegisterService = services.NewRegisterService()
	return
}
