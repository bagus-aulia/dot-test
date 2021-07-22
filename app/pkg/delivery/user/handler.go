package user

import (
	"fmt"
	"net/http"

	"github.com/bagus-aulia/dot-test/app/helpers"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	"github.com/bagus-aulia/dot-test/app/pkg/response"
	user_usecase "github.com/bagus-aulia/dot-test/app/pkg/usecase/user"
	echo "github.com/labstack/echo/v4"
)

// User represent the http handler for user
type User struct {
	Usecase user_usecase.User
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(e *echo.Echo, ucase user_usecase.User) {
	handler := &User{
		Usecase: ucase,
	}

	e.GET("/users", handler.Users)
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:uuid", handler.UserByUUID)
	e.PUT("/users/:uuid", handler.UpdateUser)
}

// Users will fetch the user list based on given params
func (h *User) Users(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.Usecase.GetUserList(ctx)
	if err != nil {
		result := response.UsersBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.UsersBody{
		Message: "SUCCESS",
		Data:    h.transformUsers(users),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// UserByUUID will fetch the user data based on given params
func (h *User) UserByUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	ctx := c.Request().Context()

	fmt.Println(uuid)

	user, err := h.Usecase.GetUserByUUID(ctx, uuid)
	if err != nil {
		result := response.UserBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.UserBody{
		Message: "SUCCESS",
		Data:    h.transformUser(&user),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// CreateUser will store the user by given request body
func (h *User) CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	ctx := c.Request().Context()

	user, err := h.Usecase.CreateUser(ctx, name, address)
	if err != nil {
		result := response.UserBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.UserBody{
		Message: "SUCCESS",
		Data:    h.transformUser(&user),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateUser will update the user by given request body
func (h *User) UpdateUser(c echo.Context) error {
	uuid := c.Param("uuid")
	name := c.FormValue("name")
	address := c.FormValue("address")
	ctx := c.Request().Context()

	user, err := h.Usecase.UpdateUser(ctx, uuid, name, address)
	if err != nil {
		result := response.UserBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.UserBody{
		Message: "SUCCESS",
		Data:    h.transformUser(&user),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

func (h *User) transformUser(user *models.User) *response.User {
	var address string = ""
	if user.Address != nil {
		address = *user.Address
	}

	res := &response.User{
		UUID:    user.UUID,
		Name:    user.Name,
		Address: address,
	}

	return res
}

func (h *User) transformUsers(users []models.User) []*response.User {
	res := make([]*response.User, 0)

	for _, user := range users {
		userRes := h.transformUser(&user)
		res = append(res, userRes)
	}

	return res
}
