package repository

import "github.com/floyoops/flo-go/pkg/contact/domain/model"

type ContactRepository interface {
	Create(contact *model.Contact) error
}
