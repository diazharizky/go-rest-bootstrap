package models

import "time"

type User struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	Password  *string    `json:"password,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
