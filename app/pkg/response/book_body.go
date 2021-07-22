package response

// BooksBody response for http json result
type BooksBody struct {
	Message string  `json:"message"`
	Error   *Error  `json:"error"`
	Data    []*Book `json:"data"`
}

// BookBody response for http json result
type BookBody struct {
	Message string `json:"message"`
	Error   *Error `json:"error"`
	Data    *Book  `json:"data"`
}

// Book is a transformer struct for Book model
type Book struct {
	Title    string `json:"title"`
	Writer   string `json:"writer"`
	UUID     string `json:"uuid"`
	Borrowed bool   `json:"borrowed"`
}
