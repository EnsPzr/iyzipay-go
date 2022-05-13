package models

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
