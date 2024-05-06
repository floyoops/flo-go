package model

import "github.com/floyoops/flo-go/pkg/core"

type Contact struct {
	Uuid      core.Identifier
	Name      string
	Email     string
	Message   string
	CreatedAt core.UtcTimeImmutable
	UpdatedAt core.UtcTimeImmutable
}

func NewContact(uuid core.Identifier, name string, email string, message string) *Contact {
	return &Contact{
		Uuid:      uuid,
		Name:      name,
		Email:     email,
		Message:   message,
		CreatedAt: core.NewUtcTimeImmutableNow(),
		UpdatedAt: core.NewUtcTimeImmutableNow(),
	}
}
