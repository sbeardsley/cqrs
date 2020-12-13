package cqrs

import (
	"fmt"
	"reflect"
)

// CommandSender is a command dispatcher
type CommandSender interface {
	Send(interface{}) error
	Register(...CommandHandler) error
}

// QuerySender is a query dispatcher
type QuerySender interface {
	SendQ(interface{}) (interface{}, error)
	RegisterQ(...QueryHandler) error
}

// InMemorySender is an in memory dispatcher
type InMemorySender struct {
	commandHandlers map[string]CommandHandler
	queryHandlers   map[string]QueryHandler
}

// NewInMemorySender constructs a new in memory dispatcher
func NewInMemorySender() *InMemorySender {
	s := &InMemorySender{
		commandHandlers: make(map[string]CommandHandler),
		queryHandlers:   make(map[string]QueryHandler),
	}
	return s
}

// Send sends a command to the command handler
func (s *InMemorySender) Send(command interface{}) error {
	cmdType := typeOf(command)
	if handler, ok := s.commandHandlers[cmdType]; ok {
		return handler.Handle(command)
	}
	return fmt.Errorf("The command bus does not have a handler for commands of type: %s", cmdType)
}

// SendQ sends a query to the query handler
func (s *InMemorySender) SendQ(query interface{}) (interface{}, error) {
	qType := typeOf(query)
	if handler, ok := s.queryHandlers[qType]; ok {
		return handler.Handle(query)
	}
	return nil, fmt.Errorf("The command bus does not have a handler for queries of type: %s", qType)
}

// Register a comamnd handler with the dispatcher
func (s *InMemorySender) Register(handlers ...CommandHandler) error {
	for _, handler := range handlers {
		cmdType := typeOf(handler.NewCommand())
		if _, ok := s.commandHandlers[cmdType]; ok {
			return fmt.Errorf("Duplicate command handler registration with command bus for command of type: %s", cmdType)
		}
		s.commandHandlers[cmdType] = handler
	}
	return nil
}

// RegisterQ a query handler with the dispatcher
func (s *InMemorySender) RegisterQ(handlers ...QueryHandler) error {
	for _, handler := range handlers {
		qType := typeOf(handler.NewQuery())
		if _, ok := s.queryHandlers[qType]; ok {
			return fmt.Errorf("Duplicate query handler registration with command bus for query of type: %s", qType)
		}
		s.queryHandlers[qType] = handler
	}
	return nil
}

func typeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}
