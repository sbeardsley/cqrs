<!-- @format -->

# cqrs simple implementation

## Installation

```
go get github.com/sbeardsley/cqrs
```

## Usage

### Command and CommandHandler

```
package command

// CreateUser command
type CreateUser struct {
	Username string
	Password string
}

// CreateUserHandler command handler
type CreateUserHandler struct {
	repo UserRepo
}

// NewCreateUserHandler command handler constructor
func NewCreateUserHandler(repo UserRepo) CreateUserHandler {
	return CreateUserHandler{
		repo: repo
	}
}

// NewCommand used during registration of handler
func (h CreateUserHandler) NewCommand() interface{} {
	return &CreateUser{}
}

// Handle command
func (h CreateUserHandler) Handle(message interface{}) error {

	cmd := message.(*CreateUser)
	err := repo.CreateUser(cmd.Username, cmd.Password)
	if err != nil {
		return err
	}

	return nil
}
```

### Initialization and registration of your handler

```
package main

import (
    "github.com/sbeardsley/cqrs"
)

func main () {
    dispatcher := cqrs.NewInMemorySender()
    userRepo := repository.NewUserRepository(db)
    cmdHandler := command.NewCreateUserHandler(userRepo)

    dispatcher.Register(cmdHandler)

    err := dispatcher.Send(&command.CreateUser{
        Username: "test_user",
        Password: "P@ssw0rd1",
    })

    if err != nil {
        panic(err)
    }
}
```

## Snippets
