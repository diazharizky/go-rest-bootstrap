package core

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/interfaces"
	"github.com/diazharizky/go-rest-bootstrap/internal/modules"
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
)

type Ctx struct {
	UserRepository    interfaces.UserRepository
	ArticleRepository interfaces.ArticleRepository

	CreateUserModule    interfaces.CreateUserModule
	CreateArticleModule interfaces.CreateArticleModule
}

func New() Ctx {
	ctx := Ctx{}

	ctx.UserRepository = repositories.NewUserRepository()
	ctx.ArticleRepository = repositories.NewArticleRepository()

	ctx.CreateUserModule = modules.NewCreateUserModule()
	ctx.CreateArticleModule = modules.NewCreateArticleModule()

	return ctx
}
