package bus

import (
	"sync"
)

type Event interface {
	Identifier() EventIdentifier
}

type EventBus interface {
	RegisterHandler(event Event, handler EventHandler)
	Publish(event Event) error
}

type EventHandler interface {
	Handle(event Event) error
}

type EventHandlerFunc func(event Event) error

func (f EventHandlerFunc) Handle(event Event) error {
	return f(event)
}

type SimpleEventBus struct {
	mu       sync.RWMutex
	handlers map[EventIdentifier][]EventHandler
}

func NewSimpleEventBus() EventBus {
	return &SimpleEventBus{
		handlers: make(map[EventIdentifier][]EventHandler),
	}
}

func (eb *SimpleEventBus) RegisterHandler(event Event, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[event.Identifier()] = append(eb.handlers[event.Identifier()], handler)
}

func (eb *SimpleEventBus) Publish(event Event) error {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	handlers, exists := eb.handlers[event.Identifier()]
	if !exists {
		return nil
	}

	for _, handler := range handlers {
		if err := handler.Handle(event); err != nil {
			return err
		}
	}

	return nil
}
