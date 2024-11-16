package bus

import "reflect"

type EventIdentifier struct {
	id string
}

func NewIdentifierFromEvent(event Event) EventIdentifier {
	return EventIdentifier{id: reflect.TypeOf(event).String()}
}

func (i EventIdentifier) String() string {
	return i.id
}
