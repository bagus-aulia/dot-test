package models

import "time"

// User Models
type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	UUID      string     `json:"uuid"`
	Address   *string    `json:"address"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
