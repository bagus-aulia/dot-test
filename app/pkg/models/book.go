package models

import (
	"time"

	"gorm.io/gorm"
)

// Book Models
type Book struct {
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	Writer    *string        `json:"writer"`
	UUID      string         `json:"uuid"`
	Borrowed  bool           `json:"borrowed" gorm:"default:false"`
	Deleted   gorm.DeletedAt `json:"deleted"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
}
