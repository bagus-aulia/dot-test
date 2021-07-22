package book

import (
	"net/http"

	"github.com/bagus-aulia/dot-test/app/helpers"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	"github.com/bagus-aulia/dot-test/app/pkg/response"
	book_usecase "github.com/bagus-aulia/dot-test/app/pkg/usecase/book"
	echo "github.com/labstack/echo/v4"
)

// Book represent the http handler for book
type Book struct {
	Usecase book_usecase.Book
}

// NewBookHandler will initialize the book/ resources endpoint
func NewBookHandler(e *echo.Echo, ucase book_usecase.Book) {
	handler := &Book{
		Usecase: ucase,
	}

	e.GET("/books", handler.Books)
	e.POST("/books", handler.CreateBook)
	e.GET("/books/:uuid", handler.BookByUUID)
	e.PUT("/books/:uuid", handler.UpdateBook)
	e.DELETE("/books/:uuid", handler.DeleteBook)
}

// Books will fetch the book list based on given params
func (h *Book) Books(c echo.Context) error {
	ctx := c.Request().Context()

	books, err := h.Usecase.GetBookList(ctx)
	if err != nil {
		result := response.BooksBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BooksBody{
		Message: "SUCCESS",
		Data:    h.transformBooks(books),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// BookByUUID will fetch the book data based on given params
func (h *Book) BookByUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	ctx := c.Request().Context()

	book, err := h.Usecase.GetBookByUUID(ctx, uuid)
	if err != nil {
		result := response.BooksBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BookBody{
		Message: "SUCCESS",
		Data:    h.transformBook(&book),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// CreateBook will store the book by given request body
func (h *Book) CreateBook(c echo.Context) error {
	title := c.FormValue("title")
	writer := c.FormValue("writer")
	ctx := c.Request().Context()

	book, err := h.Usecase.CreateBook(ctx, title, writer)
	if err != nil {
		result := response.BooksBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BookBody{
		Message: "SUCCESS",
		Data:    h.transformBook(&book),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateBook will update the book by given request body
func (h *Book) UpdateBook(c echo.Context) error {
	uuid := c.Param("uuid")
	title := c.FormValue("title")
	writer := c.FormValue("writer")
	ctx := c.Request().Context()

	book, err := h.Usecase.UpdateBook(ctx, uuid, title, writer)
	if err != nil {
		result := response.BooksBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BookBody{
		Message: "SUCCESS",
		Data:    h.transformBook(&book),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// DeleteBook will delete the book by given params
func (h *Book) DeleteBook(c echo.Context) error {
	uuid := c.Param("uuid")
	ctx := c.Request().Context()

	book, err := h.Usecase.DeleteBook(ctx, uuid)
	if err != nil {
		result := response.BooksBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BookBody{
		Message: "SUCCESS",
		Data:    h.transformBook(&book),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Book) transformBook(book *models.Book) *response.Book {
	if book == nil {
		return nil
	}

	var writer string = ""
	if book.Writer != nil {
		writer = *book.Writer
	}

	res := &response.Book{
		Title:    book.Title,
		Writer:   writer,
		UUID:     book.UUID,
		Borrowed: book.Borrowed,
	}

	return res
}

func (h *Book) transformBooks(books []models.Book) []*response.Book {
	res := make([]*response.Book, 0)

	for _, book := range books {
		bookRes := h.transformBook(&book)
		res = append(res, bookRes)
	}

	return res
}
