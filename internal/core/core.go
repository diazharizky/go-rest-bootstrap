package core

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/repositories"
	"github.com/diazharizky/go-rest-bootstrap/pkg/db"
	"github.com/diazharizky/go-rest-bootstrap/pkg/es"
)

type Core struct {
	UserRepository IUserRepository
	ItemRepository IItemRepository
}

func New() (core *Core, err error) {
	core = &Core{}

	dbConn, err := db.GetConnection()
	if err != nil {
		return
	}

	esClient, err := es.GetClient()
	if err != nil {
		return
	}

	core.UserRepository = repositories.NewUserRepository(dbConn)
	core.ItemRepository = repositories.NewItemRepository(esClient)

	return
}
