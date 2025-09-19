package services

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
)

type registerService struct {
	UserRepository repositories.UserRepository
}

func NewRegisterService() (svc registerService) {
	svc.UserRepository = repositories.NewUserRepository()
	return
}

func (svc registerService) Execute(newUser *models.User) error {
	if err := svc.UserRepository.Create(newUser); err != nil {
		return err
	}
	newUser.Password = nil

	return nil
}
