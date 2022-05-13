package models

type CheckoutFormInitializeResource struct {
	IyzipayResource
	Token               string `json:"token"`
	CheckoutFormContent string `json:"checkoutFormContent"`
	TokenExpireTime     *int64 `json:"tokenExpireTime"`
	PaymentPageUrl      string `json:"paymentPageUrl"`
}
