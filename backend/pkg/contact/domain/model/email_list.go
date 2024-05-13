package model

type EmailList struct {
	Emails []*Email
}

func NewEmailList(emails []*Email) *EmailList {
	return &EmailList{emails}
}

func (el *EmailList) ToArray() []*Email {
	return el.Emails
}

func (el *EmailList) ToArrayString() []string {
	var emailsStrings []string
	for _, email := range el.Emails {
		emailsStrings = append(emailsStrings, email.String())
	}
	return emailsStrings
}
