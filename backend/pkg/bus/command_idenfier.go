package bus

import (
	"reflect"
)

type CommandIdentifier struct {
	id string
}

func NewIdentifierFromCommand(command Command) CommandIdentifier {
	return CommandIdentifier{id: reflect.TypeOf(command).String()}
}

func (i CommandIdentifier) String() string {
	return i.id
}
