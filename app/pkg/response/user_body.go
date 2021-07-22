package response

// UsersBody response for http json result
type UsersBody struct {
	Message string  `json:"message"`
	Error   *Error  `json:"error"`
	Data    []*User `json:"data"`
}

// UserBody response for http json result
type UserBody struct {
	Message string `json:"message"`
	Error   *Error `json:"error"`
	Data    *User  `json:"data"`
}

// User is a transformer struct for User model
type User struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Address string `json:"address"`
}
