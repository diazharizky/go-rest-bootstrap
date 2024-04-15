package models

import "time"

type Article struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	AuthorID  int64      `json:"author_id" gorm:"type:bigint"`
	Title     string     `json:"title" gorm:"size:100;not null"`
	Content   string     `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null;default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Author    User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
