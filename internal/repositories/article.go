package repositories

import (
	"github.com/diazharizky/go-rest-bootstrap/internal/models"
	"github.com/diazharizky/go-rest-bootstrap/pkg/db"
)

type articleRepository struct {
	tableName string
}

func NewArticleRepository() (r articleRepository) {
	r.tableName = "articles"
	return
}

func (r articleRepository) List(userID int64) ([]models.Article, error) {
	conn := db.MustGetConnection()

	articles := []models.Article{}
	if tx := conn.Where("author_id = ?", userID).Find(&articles); tx.Error != nil {
		return nil, tx.Error
	}

	return articles, nil
}

func (r articleRepository) Get(articleID int64) (*models.Article, error) {
	conn := db.MustGetConnection()

	article := models.Article{}
	if tx := conn.First(&article, articleID); tx.Error != nil {
		return nil, tx.Error
	}

	return &article, nil
}

func (r articleRepository) Create(newArticle *models.Article) error {
	conn := db.MustGetConnection()

	if tx := conn.Create(&newArticle); tx.Error != nil {
		return tx.Error
	}

	return nil
}
