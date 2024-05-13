package core

import (
	"errors"
	"github.com/google/uuid"
)

type Identifier struct {
	id string
}

func NewIdentifier() Identifier {
	id := uuid.New().String()
	return Identifier{id: id}
}

func NewIdentifierFromString(idStr string) (Identifier, error) {
	_, err := uuid.Parse(idStr)
	if err != nil {
		return Identifier{}, errors.New("Invalid UUID")
	}
	return Identifier{id: idStr}, nil
}

func (id Identifier) String() string {
	return id.id
}
