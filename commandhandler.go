package cqrs

// CommandHandler handles commands
type CommandHandler interface {
	Handle(interface{}) error
	NewCommand() interface{}
}

// QueryHandler handles queries
type QueryHandler interface {
	Handle(interface{}) (interface{}, error)
	NewQuery() interface{}
}
