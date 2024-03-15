package models

import "time"

type Article struct {
	ID        int32      `json:"id"`
	AuthorID  int32      `json:"author_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
