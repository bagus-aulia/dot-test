package models

import "time"

// BorrowDetail Models
type BorrowDetail struct {
	ID        uint       `json:"id"`
	BorrowID  uint       `json:"borrow_id"`
	BookID    *uint      `json:"book_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Book *Book
}
