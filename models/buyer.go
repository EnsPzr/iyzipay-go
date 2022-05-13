package models

import "github.com/enspzr/iyzipay-go/helpers"

type Buyer struct {
	Id                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Surname             string `json:"surname,omitempty"`
	IdentityNumber      string `json:"identityNumber,omitempty"`
	Email               string `json:"email,omitempty"`
	GsmNumber           string `json:"gsmNumber,omitempty"`
	RegistrationDate    string `json:"registrationDate,omitempty"`
	LastLoginDate       string `json:"lastLoginDate,omitempty"`
	RegistrationAddress string `json:"registrationAddress,omitempty"`
	City                string `json:"city,omitempty"`
	Country             string `json:"country,omitempty"`
	ZipCode             string `json:"zipCode,omitempty"`
	Ip                  string `json:"ip,omitempty"`
}

func (b *Buyer) String() string {
	t := ""
	t += helpers.Append("id", b.Id)
	t += helpers.Append("name", b.Name)
	t += helpers.Append("surname", b.Surname)
	t += helpers.Append("identityNumber", b.IdentityNumber)
	t += helpers.Append("email", b.Email)
	t += helpers.Append("gsmNumber", b.GsmNumber)
	t += helpers.Append("registrationDate", b.RegistrationDate)
	t += helpers.Append("lastLoginDate", b.LastLoginDate)
	t += helpers.Append("registrationAddress", b.RegistrationAddress)
	t += helpers.Append("city", b.City)
	t += helpers.Append("country", b.Country)
	t += helpers.Append("zipCode", b.ZipCode)
	t += helpers.Append("ip", b.Ip)
	return t
}
