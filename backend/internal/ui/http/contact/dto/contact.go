package dto

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

const (
	required string = "required"
	email    string = "email"
	mini     string = "min"
	maxi     string = "max"
)

const (
	ValidationErrContactName    = "Please enter a name with 3 to 50 characters."
	ValidationErrContactEmail   = "Please enter a email."
	ValidationErrContactMessage = "Please enter a name with 3 to 1000 characters."
)

type ContactDto struct {
	Name    string `json:"name" validate:"required,min=3,max=50"`
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required,min=3,max=1000"`
}

func NewContactDto() *ContactDto {
	return &ContactDto{}
}

func FromBody(body io.ReadCloser) (*ContactDto, error) {
	dto := NewContactDto()
	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (c *ContactDto) Validate() map[string]string {
	return validateDto(c)
}

func validateDto(c interface{}) map[string]string {
	err := validator.New().Struct(c)
	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)
	if len(errors) == 0 {
		return nil
	}

	return createErrorMessages(errors)
}

func createErrorMessages(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)
	for i := range errors {
		switch errors[i].StructField() {
		case "Name":
			switch errors[i].Tag() {
			case required, mini, maxi:
				result["name"] = ValidationErrContactName
			}
		case "Email":
			switch errors[i].Tag() {
			case required, email:
				result["email"] = ValidationErrContactEmail
			}
		case "Message":
			switch errors[i].Tag() {
			case required, mini, maxi:
				result["message"] = ValidationErrContactMessage
			}
		}

	}
	return result
}
