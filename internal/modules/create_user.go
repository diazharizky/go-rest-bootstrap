package modules

import (
	"fmt"

	"github.com/diazharizky/go-rest-bootstrap/internal/interfaces"
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
	"github.com/diazharizky/go-rest-bootstrap/pkg/emailclient"
)

type createUserModule struct {
	userRepository interfaces.UserRepository
}

func NewCreateUserModule() (m createUserModule) {
	m.userRepository = repositories.NewUserRepository()
	return
}

func (m createUserModule) Call(newUser *models.User) error {
	if err := m.userRepository.Create(newUser); err != nil {
		return err
	}

	emailClient := emailclient.New()

	if err := emailClient.SendNoAuth([]string{newUser.Email}, nil, "Registration Success", "Congratulations!"); err != nil {
		fmt.Println(
			fmt.Errorf("error while sending email: %v", err),
		)
	}

	return nil
}
