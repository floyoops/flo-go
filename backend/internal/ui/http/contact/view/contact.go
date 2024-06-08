package view

import "github.com/floyoops/flo-go/backend/internal/ui/http/contact/dto"

type ContactView struct {
	Data   *dto.ContactDto
	Errors *map[string]string
	Sended bool
}

func NewContactView(Data *dto.ContactDto, Errors *map[string]string, Sended bool) ContactView {
	return ContactView{Data: Data, Errors: Errors, Sended: Sended}
}
