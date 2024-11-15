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

type CommandBus struct {
	handlers map[CommandIdentifier]CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[CommandIdentifier]CommandHandler),
	}
}

func (cb *CommandBus) RegisterHandler(command Command, handler CommandHandler) {
	cb.handlers[command.Identifier()] = handler
}

func (cb *CommandBus) Dispatch(command Command) error {
	handler, ok := cb.handlers[command.Identifier()]
	if !ok {
		return errors.New("no handler registered for this command")

	}
	return handler.Handle(command)
}
