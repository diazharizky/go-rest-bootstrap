package services

import (
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
	"github.com/diazharizky/go-rest-bootstrap/pkg/emailclient"
	"github.com/diazharizky/go-rest-bootstrap/templates"
)

type createUserService struct {
	UserRepository repositories.UserRepository
}

func NewCreateUserService() (m createUserService) {
	m.UserRepository = repositories.NewUserRepository()
	return
}

func (m createUserService) Call(newUser *models.User) error {
	if err := m.UserRepository.Create(newUser); err != nil {
		return err
	}

	emailClient := emailclient.New()
	tpl := templates.RegistrationConfirmationTemplate{
		Name:  "Some Cool Name",
		Email: newUser.Email,
	}
	if err := emailClient.SendNoAuth([]string{newUser.Email}, nil, "Registration Success", &tpl); err != nil {
		fmt.Println(
			fmt.Errorf("error while sending email: %v", err),
		)
	}

	return nil
}
