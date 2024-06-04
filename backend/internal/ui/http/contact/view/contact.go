package view

import (
	. "github.com/floyoops/flo-go/internal/ui/http/contact/dto"
)

type ContactView struct {
	Data   *ContactDto
	Errors *map[string]string
	Sended bool
}

func NewContactView(Data *ContactDto, Errors *map[string]string, Sended bool) ContactView {
	return ContactView{Data: Data, Errors: Errors, Sended: Sended}
}
