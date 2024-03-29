package models

import "time"

type User struct {
	ID        int32      `json:"id"`
	Email     string     `json:"email"`
	FullName  string     `json:"full_name"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
