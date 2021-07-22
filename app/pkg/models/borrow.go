package models

import (
	"time"
)

// Borrow Models
type Borrow struct {
	ID        uint       `json:"id"`
	UUID      string     `json:"uuid"`
	UserID    *uint      `json:"user_id"`
	StartAt   time.Time  `json:"start_at"`
	EndAt     *time.Time `json:"end_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	User         *User
	BorrowDetail []BorrowDetail
}
