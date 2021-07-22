package borrow_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	borrowDelivery "github.com/bagus-aulia/dot-test/app/pkg/delivery/borrow"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	borrowUsecaseMock "github.com/bagus-aulia/dot-test/app/pkg/usecase/borrow/mocks"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBorrows(t *testing.T) {
	var mockBorrow models.Borrow
	err := faker.FakeData(&mockBorrow)
	assert.NoError(t, err)

	mockBorrowUcase := new(borrowUsecaseMock.BorrowUsecase)
	mockBorrowList := make([]models.Borrow, 0)
	mockBorrowList = append(mockBorrowList, mockBorrow)

	mockBorrowUcase.On("GetBorrowList", mock.Anything).Return(mockBorrowList, nil)
	defer mockBorrowUcase.On("GetBorrowList", mock.Anything)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/borrows", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := borrowDelivery.Borrow{
		Usecase: mockBorrowUcase,
	}
	err = handler.Borrows(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBorrowUcase.AssertExpectations(t)
}

func TestBorrowByUUID(t *testing.T) {
	var mockBorrow models.Borrow
	err := faker.FakeData(&mockBorrow)
	assert.NoError(t, err)

	mockBorrowUcase := new(borrowUsecaseMock.BorrowUsecase)
	uuid := mockBorrow.UUID

	mockBorrowUcase.On("GetBorrowByUUID", mock.Anything, uuid).Return(mockBorrow, nil)
	defer mockBorrowUcase.On("GetBorrowByUUID", mock.Anything, uuid)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/borrows/"+uuid, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("borrows/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	handler := borrowDelivery.Borrow{
		Usecase: mockBorrowUcase,
	}
	err = handler.BorrowByUUID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBorrowUcase.AssertExpectations(t)
}

func TestCreateBorrow(t *testing.T) {
	var mockBorrow models.Borrow
	err := faker.FakeData(&mockBorrow)
	assert.NoError(t, err)

	mockBorrowUcase := new(borrowUsecaseMock.BorrowUsecase)

	mockBorrowUcase.On("CreateBorrow", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(mockBorrow, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/borrows", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/borrows")
	c.SetParamNames("user_uuid")
	c.SetParamValues("BpLnf-gDsc2WD-8F2qUNfH-210722065900")
	c.SetParamNames("book_uuids[]")
	c.SetParamValues("K5a84-jjJkwzD-kh9hE2fh-210722131208")

	handler := borrowDelivery.Borrow{
		Usecase: mockBorrowUcase,
	}
	err = handler.CreateBorrow(c)
	require.NoError(t, err)

	assert.Equal(t, 500, rec.Code)
}

func TestReturnBorrow(t *testing.T) {
	var mockBorrow models.Borrow
	err := faker.FakeData(&mockBorrow)
	assert.NoError(t, err)

	mockBorrowUcase := new(borrowUsecaseMock.BorrowUsecase)
	uuid := mockBorrow.UUID

	mockBorrowUcase.On("ReturnBorrow", mock.Anything, mock.AnythingOfType("string")).Return(mockBorrow, nil)
	defer mockBorrowUcase.On("ReturnBorrow", mock.Anything, mock.AnythingOfType("string"))

	e := echo.New()
	req, err := http.NewRequest(echo.PATCH, "/borrows/"+uuid, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/borrows/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	handler := borrowDelivery.Borrow{
		Usecase: mockBorrowUcase,
	}
	err = handler.ReturnBorrow(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBorrowUcase.AssertExpectations(t)
}
