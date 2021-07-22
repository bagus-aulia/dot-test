package response

import (
	"time"
)

// BorrowsBody response for http json result
type BorrowsBody struct {
	Message string    `json:"message"`
	Error   *Error    `json:"error"`
	Data    []*Borrow `json:"data"`
}

// BorrowBody response for http json result
type BorrowBody struct {
	Message string  `json:"message"`
	Error   *Error  `json:"error"`
	Data    *Borrow `json:"data"`
}

// Borrow is a transformer struct for Borrow model
type Borrow struct {
	UUID    string     `json:"uuid"`
	StartAt time.Time  `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`

	User         *User           `json:"user"`
	BorrowDetail []*BorrowDetail `json:"items"`
}

// BorrowDetail is a transformer struct for BorrowDetail model
type BorrowDetail struct {
	Book *Book `json:"book"`
}
