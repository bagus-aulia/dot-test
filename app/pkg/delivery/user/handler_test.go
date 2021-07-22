package user_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	userDelivery "github.com/bagus-aulia/dot-test/app/pkg/delivery/user"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	userUsecaseMock "github.com/bagus-aulia/dot-test/app/pkg/usecase/user/mocks"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUserUcase := new(userUsecaseMock.UserUsecase)
	mockUserList := make([]models.User, 0)
	mockUserList = append(mockUserList, mockUser)

	mockUserUcase.On("GetUserList", mock.Anything).Return(mockUserList, nil)
	defer mockUserUcase.On("GetUserList", mock.Anything)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := userDelivery.User{
		Usecase: mockUserUcase,
	}
	err = handler.Users(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserUcase.AssertExpectations(t)
}

func TestUserByUUID(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUserUcase := new(userUsecaseMock.UserUsecase)
	uuid := mockUser.UUID

	mockUserUcase.On("GetUserByUUID", mock.Anything, uuid).Return(mockUser, nil)
	defer mockUserUcase.On("GetUserByUUID", mock.Anything, uuid)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users/"+uuid, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("users/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	handler := userDelivery.User{
		Usecase: mockUserUcase,
	}
	err = handler.UserByUUID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserUcase.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUserUcase := new(userUsecaseMock.UserUsecase)

	mockUserUcase.On("CreateUser", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockUser, nil)
	defer mockUserUcase.On("CreateUser", mock.Anything, mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"))

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")
	c.SetParamNames("name")
	c.SetParamValues("Gerald Newt")
	c.SetParamNames("address")
	c.SetParamValues("Laguna Lagoon")

	handler := userDelivery.User{
		Usecase: mockUserUcase,
	}
	err = handler.CreateUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserUcase.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUserUcase := new(userUsecaseMock.UserUsecase)
	uuid := mockUser.UUID

	mockUserUcase.On("UpdateUser", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockUser, nil)
	defer mockUserUcase.On("UpdateUser", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"))

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/users/"+uuid, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("users/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)
	c.SetParamNames("name")
	c.SetParamValues("Gerald Newt")
	c.SetParamNames("address")
	c.SetParamValues("Laguna Lagoon")

	handler := userDelivery.User{
		Usecase: mockUserUcase,
	}
	err = handler.UpdateUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserUcase.AssertExpectations(t)
}
