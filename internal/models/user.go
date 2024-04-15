package models

import "time"

type User struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email" gorm:"size:30;unique;not null"`
	FullName  string     `json:"full_name" gorm:"size:30;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
