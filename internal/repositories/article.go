package repositories

import (
	"database/sql"
	"fmt"

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
	db := db.GetConnection()
	defer db.Close()

	rows, err := db.Query(
		fmt.Sprintf("SELECT id, author_id, title, content, created_at, updated_at FROM %s WHERE author_id=%d", r.tableName, userID),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []models.Article{}

	article := models.Article{}
	for rows.Next() {
		err := rows.Scan(&article.ID, &article.AuthorID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (r articleRepository) Get(articleID int64) (*models.Article, error) {
	db := db.GetConnection()
	defer db.Close()

	row := db.QueryRow(
		fmt.Sprintf("SELECT id, author_id, title, content, created_at, deleted_at FROM %s WHERE id=$1", r.tableName),
		articleID,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	article := models.Article{}

	err := row.Scan(&article.ID, &article.AuthorID, &article.Title, &article.Content, &article.CreatedAt, &article.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &article, nil
}

func (r articleRepository) Create(newArticle *models.Article) error {
	db := db.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(
		fmt.Sprintf("INSERT INTO %s(author_id, title, content) VALUES($1, $2, $3) RETURNING id, created_at, updated_at", r.tableName),
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.
		QueryRow(
			newArticle.AuthorID, newArticle.Title, newArticle.Content,
		).
		Scan(
			&newArticle.ID, &newArticle.CreatedAt, &newArticle.UpdatedAt,
		)

	if err != nil {
		return err
	}

	return nil
}
