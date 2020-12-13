package main

import (
	"fmt"

	"github.com/sbeardsley/cqrs/example/query"

	"github.com/sbeardsley/cqrs"

	"github.com/sbeardsley/cqrs/example/command"
)

func main() {
	fmt.Println("starting example")

	dispatcher := cqrs.NewInMemorySender()
	cmdHandler := command.NewCreateUserHandler()
	qryHandler := query.NewGetUserHandler()

	dispatcher.Register(cmdHandler)
	dispatcher.RegisterQ(qryHandler)

	newUser := &command.CreateUser{Username: "sbeardsley", Password: "P@ssw0rd1"}

	err := dispatcher.Send(newUser)
	if err != nil {
		fmt.Println(err)
	}

	usr, err := dispatcher.SendQ(&query.GetUser{ID: 1})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("got user %v", usr)
}
