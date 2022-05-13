package models

type CheckoutFormInitialize struct {
	CheckoutFormInitializeResource
}

func (c CheckoutFormInitialize) IsSuccess() bool {
	if c.Status == "success" {
		return true
	}
	return false
}
