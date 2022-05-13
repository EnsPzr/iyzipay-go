package models

import "github.com/enspzr/iyzipay-go/helpers"

type Address struct {
	Description string `json:"address,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
	ContactName string `json:"contactName,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
}

func (a *Address) String() string {
	t := ""
	t += helpers.Append("address", a.Description)
	t += helpers.Append("zipCode", a.ZipCode)
	t += helpers.Append("contactName", a.ContactName)
	t += helpers.Append("city", a.City)
	t += helpers.Append("country", a.Country)
	return t
}
