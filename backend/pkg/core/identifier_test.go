package core

import (
	"testing"
)

func TestNewIdentifier(t *testing.T) {
	id := NewIdentifier()
	if id.String() == "" {
		t.Errorf("NewIdentifier() returned an empty ID")
	}
}

func TestNewIdentifierFromString(t *testing.T) {
	validUUID := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	invalidUUID := "not-a-valid-uuid"

	id, err := NewIdentifierFromString(validUUID)
	if err != nil {
		t.Errorf("Unexpected error creating Identifier from valid UUID string: %v", err)
	}
	if id.String() != validUUID {
		t.Errorf("Identifier created from valid UUID string does not match")
	}

	_, err = NewIdentifierFromString(invalidUUID)
	if err == nil {
		t.Errorf("No error returned when creating Identifier from invalid UUID string")
	}
}
