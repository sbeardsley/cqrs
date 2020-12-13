package command

import (
	"fmt"
)

// CreateUser command
type CreateUser struct {
	Username string
	Password string
}

// CreateUserHandler command handler
type CreateUserHandler struct {
}

// NewCreateUserHandler command handler constructor
func NewCreateUserHandler() CreateUserHandler {
	return CreateUserHandler{}
}

// NewCommand used during registration of handler
func (h CreateUserHandler) NewCommand() interface{} {
	return &CreateUser{}
}

// Handle command
func (h CreateUserHandler) Handle(message interface{}) error {

	cmd := message.(*CreateUser)
	fmt.Printf("Executing create user handler %v", cmd)

	//Create the user

	return nil
}
