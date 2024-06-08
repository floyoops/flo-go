package view

import (
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact/dto"
)

type ContactView struct {
	Data   *dto.ContactDto    `json:"data"`
	Errors *map[string]string `json:"errors"`
	Sent   bool               `json:"sent"`
}

func NewContactView(Data *dto.ContactDto, Errors *map[string]string, Sent bool) ContactView {
	return ContactView{Data: Data, Errors: Errors, Sent: Sent}
}
