package book_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	bookDelivery "github.com/bagus-aulia/dot-test/app/pkg/delivery/book"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	bookUsecaseMock "github.com/bagus-aulia/dot-test/app/pkg/usecase/book/mocks"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBooks(t *testing.T) {
	var mockBook models.Book
	err := faker.FakeData(&mockBook)
	assert.NoError(t, err)

	mockBookUcase := new(bookUsecaseMock.BookUsecase)
	mockBookList := make([]models.Book, 0)
	mockBookList = append(mockBookList, mockBook)

	mockBookUcase.On("GetBookList", mock.Anything).Return(mockBookList, nil)
	defer mockBookUcase.On("GetBookList", mock.Anything)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/books", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := bookDelivery.Book{
		Usecase: mockBookUcase,
	}
	err = handler.Books(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBookUcase.AssertExpectations(t)
}

func TestBookByUUID(t *testing.T) {
	var mockBook models.Book
	err := faker.FakeData(&mockBook)
	assert.NoError(t, err)

	mockBookUcase := new(bookUsecaseMock.BookUsecase)
	uuid := mockBook.UUID

	mockBookUcase.On("GetBookByUUID", mock.Anything, uuid).Return(mockBook, nil)
	defer mockBookUcase.On("GetBookByUUID", mock.Anything, uuid)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/books/"+uuid, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	handler := bookDelivery.Book{
		Usecase: mockBookUcase,
	}
	err = handler.BookByUUID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBookUcase.AssertExpectations(t)
}

func TestCreateBook(t *testing.T) {
	var mockBook models.Book
	err := faker.FakeData(&mockBook)
	assert.NoError(t, err)

	mockBookUcase := new(bookUsecaseMock.BookUsecase)

	mockBookUcase.On("CreateBook", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockBook, nil)
	defer mockBookUcase.On("CreateBook", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"))

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/books", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books")
	c.SetParamNames("title")
	c.SetParamValues("That time that I reincarnated")
	c.SetParamNames("writer")
	c.SetParamValues("Masashi Uramoto")

	handler := bookDelivery.Book{
		Usecase: mockBookUcase,
	}
	err = handler.CreateBook(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBookUcase.AssertExpectations(t)
}

func TestUpdateBook(t *testing.T) {
	var mockBook models.Book
	err := faker.FakeData(&mockBook)
	assert.NoError(t, err)

	mockBookUcase := new(bookUsecaseMock.BookUsecase)
	uuid := mockBook.UUID

	mockBookUcase.On("UpdateBook", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockBook, nil)
	defer mockBookUcase.On("UpdateBook", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"))

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/books/"+uuid, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)
	c.SetParamNames("title")
	c.SetParamValues("That time that I reincarnated")
	c.SetParamNames("writer")
	c.SetParamValues("Masashi Uramoto")

	handler := bookDelivery.Book{
		Usecase: mockBookUcase,
	}
	err = handler.UpdateBook(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBookUcase.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	var mockBook models.Book
	err := faker.FakeData(&mockBook)
	assert.NoError(t, err)

	mockBookUcase := new(bookUsecaseMock.BookUsecase)
	uuid := mockBook.UUID

	mockBookUcase.On("DeleteBook", mock.Anything, uuid).Return(mockBook, nil)
	defer mockBookUcase.On("DeleteBook", mock.Anything, uuid)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/books/"+uuid, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	handler := bookDelivery.Book{
		Usecase: mockBookUcase,
	}
	err = handler.DeleteBook(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBookUcase.AssertExpectations(t)
}
