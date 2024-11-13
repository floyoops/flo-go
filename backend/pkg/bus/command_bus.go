package bus

import (
	"errors"
)

type Command interface {
	Identifier() CommandIdentifier
}

type Handler interface {
	Handle(Command) error
}

type CommandBus struct {
	handlers map[CommandIdentifier]Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[CommandIdentifier]Handler),
	}
}

func (cb *CommandBus) RegisterHandler(command Command, handler Handler) {
	cb.handlers[command.Identifier()] = handler
}

func (cb *CommandBus) Dispatch(command Command) error {
	handler, ok := cb.handlers[command.Identifier()]
	if !ok {
		return errors.New("no handler registered for this command")

	}
	return handler.Handle(command)
}
