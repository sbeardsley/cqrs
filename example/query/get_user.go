package query

import "fmt"

// GetUser query
type GetUser struct {
	ID int
}

// User model
type User struct {
	Username string
}

// GetUserHandler query handler
type GetUserHandler struct {
}

// NewGetUserHandler query handler constructor
func NewGetUserHandler() GetUserHandler {
	return GetUserHandler{}
}

// NewQuery used during registration of handler
func (h GetUserHandler) NewQuery() interface{} {
	return &GetUser{}
}

// Handle command
func (h GetUserHandler) Handle(message interface{}) (interface{}, error) {

	qry := message.(*GetUser)
	fmt.Printf("Executing get user handler %v", qry)

	return &User{
		Username: "test_user",
	}, nil
}
