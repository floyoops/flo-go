package repository

import (
	"errors"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
)

type ContactRepository interface {
	Create(contact *model.Contact) error
}

var (
	ErrOnSaveContact = errors.New("error on save contact")
)
