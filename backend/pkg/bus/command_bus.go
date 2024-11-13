package bus

import (
	"errors"
	"reflect"
)

type Command interface{}

type Handler interface {
	Handle(Command) error
}

type CommandBus struct {
	handlers map[reflect.Type]Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[reflect.Type]Handler),
	}
}

func (cb *CommandBus) RegisterHandler(command Command, handler Handler) {
	commandType := reflect.TypeOf(command)
	cb.handlers[commandType] = handler
}

func (cb *CommandBus) Dispatch(command Command) error {
	commandType := reflect.TypeOf(command)
	handler, ok := cb.handlers[commandType]
	if !ok {
		return errors.New("no handler registered for this command")

	}
	return handler.Handle(command)
}
