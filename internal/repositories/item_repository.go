package repositories

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/elastic/go-elasticsearch/v8"
)

type itemRepository struct {
	es *elasticsearch.Client
}

func NewItemRepository(es *elasticsearch.Client) itemRepository {
	return itemRepository{
		es: es,
	}
}

func (rep itemRepository) List() ([]models.Item, error) {
	return []models.Item{}, nil
}
