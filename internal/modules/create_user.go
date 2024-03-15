package modules

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/interfaces"
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
)

type createUserModule struct {
	userRepository interfaces.UserRepository
}

func NewCreateUserModule() (m createUserModule) {
	m.userRepository = repositories.NewUserRepository()
	return
}

func (m createUserModule) Call(newUser *models.User) error {
	return m.userRepository.Create(newUser)
}
