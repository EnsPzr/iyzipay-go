package models

type Address struct {
	Description string `json:"address,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
	ContactName string `json:"contactName,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
}
