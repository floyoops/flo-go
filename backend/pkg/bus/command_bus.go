package bus

import (
	"errors"
)

type Command interface {
	Identifier() CommandIdentifier
}

type CommandHandler interface {
	Handle(Command) error
}

type CommandMiddleware func(next CommandHandler) CommandHandler

type CommandBus struct {
	handlers    map[CommandIdentifier]CommandHandler
	middlewares []CommandMiddleware
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[CommandIdentifier]CommandHandler),
	}
}

func (cb *CommandBus) RegisterHandler(command Command, handler CommandHandler) {
	cb.handlers[command.Identifier()] = handler
}

func (cb *CommandBus) Use(middleware CommandMiddleware) {
	cb.middlewares = append(cb.middlewares, middleware)
}

func (cb *CommandBus) Dispatch(command Command) error {
	handler, ok := cb.handlers[command.Identifier()]
	if !ok {
		return errors.New("no handler registered for this command")

	}

	// Apply middlewares
	finalHandler := handler
	for i := len(cb.middlewares) - 1; i >= 0; i-- {
		finalHandler = cb.middlewares[i](finalHandler)
	}

	return finalHandler.Handle(command)
}

/*
CommandHandlerFunc is a functional type that implements CommandHandler.
Used for middlewares.
*/
type CommandHandlerFunc func(command Command) error

func (f CommandHandlerFunc) Handle(command Command) error {
	return f(command)
}
