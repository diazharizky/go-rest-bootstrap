package models

import "time"

type Article struct {
	ID        int64      `json:"id"`
	AuthorID  int64      `json:"authorId"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Author    User       `json:"author,omitempty"`
}
