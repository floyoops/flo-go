package bus

import (
	"errors"
)

type Command interface {
	Identifier() CommandIdentifier
}

type CommandHandler interface {
	Handle(Command) ([]Event, error)
}

type CommandMiddleware func(next CommandHandler) CommandHandler

type CommandBus struct {
	handlers    map[CommandIdentifier]CommandHandler
	middlewares []CommandMiddleware
	eventBus    *EventBus
}

func NewCommandBus(eventBus *EventBus) *CommandBus {
	return &CommandBus{
		handlers: make(map[CommandIdentifier]CommandHandler),
		eventBus: eventBus,
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

	// Publish events
	events, err := finalHandler.Handle(command)
	if err != nil {
		return err
	}

	for _, event := range events {
		if err := cb.eventBus.Publish(event); err != nil {
			return err
		}
	}

	return nil
}

/*
CommandHandlerFunc is a functional type that implements CommandHandler.
Used for middlewares.
*/
type CommandHandlerFunc func(command Command) ([]Event, error)

func (f CommandHandlerFunc) Handle(command Command) ([]Event, error) {
	return f(command)
}
