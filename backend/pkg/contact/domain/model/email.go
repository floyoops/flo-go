package model

import (
	"errors"
	"fmt"
	"regexp"
)

type Email struct {
	address string
}

func NewEmail(address string) (*Email, error) {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, address)
	if !match {
		return nil, errors.New(fmt.Sprintf("'%s' is not a valid email format", address))
	}
	return &Email{address: address}, nil
}

func (e *Email) String() string {
	return e.address
}
