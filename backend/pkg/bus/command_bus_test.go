package bus

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Mocks
 */
type MockCommand struct {
	name string
}

func (mc MockCommand) Identifier() CommandIdentifier {
	return CommandIdentifier{id: mc.name}
}

type MockEvent struct {
	name string
}

func (e MockEvent) Identifier() EventIdentifier {
	return EventIdentifier{id: e.name}
}

type MockEventBus struct {
	PublishedEvents []Event
}

func (eb *MockEventBus) RegisterHandler(event Event, handler EventHandler) {
}
func (eb *MockEventBus) Publish(event Event) error {
	eb.PublishedEvents = append(eb.PublishedEvents, event)
	return nil
}

type MockFailingEventBus struct{}

func (eb *MockFailingEventBus) RegisterHandler(event Event, handler EventHandler) {
}
func (eb *MockFailingEventBus) Publish(event Event) error {
	return errors.New("event publishing failed")
}

/**
 * Tests
 */
func TestCommandBus_ShouldPublishOneEventOnDispatchCommand(t *testing.T) {
	// Given
	eventBus := &MockEventBus{}
	commandBus := NewCommandBus(eventBus)

	mockCommand := MockCommand{name: "mock-command"}
	mockHandler := CommandHandlerFunc(func(command Command) ([]Event, error) {
		return []Event{MockEvent{name: "event1"}}, nil
	})
	commandBus.RegisterHandler(mockCommand, mockHandler)

	// When
	err := commandBus.Dispatch(mockCommand)

	// Then
	assert.NoError(t, err)
	assert.Len(t, eventBus.PublishedEvents, 1)
	assert.Equal(t, "event1", eventBus.PublishedEvents[0].(MockEvent).name)
}

func TestCommandBus_ShouldReturnErrorOnDispatchWithNoRegisteredHandler(t *testing.T) {
	// Given
	eventBus := &MockEventBus{}
	commandBus := NewCommandBus(eventBus)
	mockCommand := MockCommand{name: "mock-command"}

	// When
	err := commandBus.Dispatch(mockCommand)

	// Then
	assert.Error(t, err)
	assert.Equal(t, "no handler registered for this command", err.Error())
}

func TestCommandBus_ShouldExecuteTwoMiddlewares(t *testing.T) {
	// Given
	eventBus := &MockEventBus{}
	commandBus := NewCommandBus(eventBus)
	mockCommand := MockCommand{name: "mock-command"}
	mockHandler := CommandHandlerFunc(func(command Command) ([]Event, error) {
		return []Event{MockEvent{name: "event1"}}, nil
	})

	// Middleware to track execution order
	executionOrder := []string{}
	middleware1 := func(next CommandHandler) CommandHandler {
		return CommandHandlerFunc(func(command Command) ([]Event, error) {
			executionOrder = append(executionOrder, "middleware1")
			return next.Handle(command)
		})
	}

	middleware2 := func(next CommandHandler) CommandHandler {
		return CommandHandlerFunc(func(command Command) ([]Event, error) {
			executionOrder = append(executionOrder, "middleware2")
			return next.Handle(command)
		})
	}

	commandBus.RegisterHandler(mockCommand, mockHandler)
	commandBus.Use(middleware1)
	commandBus.Use(middleware2)

	// When
	err := commandBus.Dispatch(mockCommand)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, []string{"middleware1", "middleware2"}, executionOrder)
	assert.Len(t, eventBus.PublishedEvents, 1)
	assert.Equal(t, "event1", eventBus.PublishedEvents[0].(MockEvent).name)
}

func TestCommandBus_ShouldReturnAnHandlerError(t *testing.T) {
	// Given
	eventBus := &MockEventBus{}
	commandBus := NewCommandBus(eventBus)

	mockCommand := MockCommand{name: "mock-command"}
	mockHandler := CommandHandlerFunc(func(command Command) ([]Event, error) {
		return nil, errors.New("handler error")
	})

	commandBus.RegisterHandler(mockCommand, mockHandler)

	// When
	err := commandBus.Dispatch(mockCommand)

	// Then
	assert.Error(t, err)
	assert.Equal(t, "handler error", err.Error())
	assert.Empty(t, eventBus.PublishedEvents)
}

func TestCommandBus_EventBusError(t *testing.T) {
	// Given
	eventBus := &MockFailingEventBus{}
	commandBus := NewCommandBus(eventBus)

	mockCommand := MockCommand{name: "mock-command"}
	mockHandler := CommandHandlerFunc(func(command Command) ([]Event, error) {
		return []Event{MockEvent{name: "event1"}}, nil
	})

	commandBus.RegisterHandler(mockCommand, mockHandler)

	// When
	err := commandBus.Dispatch(mockCommand)

	// Then
	assert.Error(t, err)
	assert.Equal(t, "event publishing failed", err.Error())
}
