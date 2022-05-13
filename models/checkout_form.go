package models

type CheckoutForm struct {
	Token                        string  `json:"token"`
	CallbackUrl                  string  `json:"callbackUrl"`
	ErrorMessage                 string  `json:"errorMessage"`
	ErrorGroup                   string  `json:"errorGroup"`
	ErrorCode                    string  `json:"errorStatus"`
	Price                        float32 `json:"price"`
	PaidPrice                    float32 `json:"paidPrice"`
	Installment                  *int    `json:"installment"`
	Currency                     string  `json:"currency"`
	PaymentId                    string  `json:"paymentId"`
	PaymentStatus                string  `json:"paymentStatus"`
	FraudStatus                  *int    `json:"fraudStatus"`
	MerchantCommissionRate       float32 `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float32 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float32 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float32 `json:"iyziCommissionFee"`
	CardType                     string  `json:"cardType"`
	CardAssociation              string  `json:"cardAssociation"`
	CardFamily                   string  `json:"cardFamily"`
	CardToken                    string  `json:"cardToken"`
	CardUserKey                  string  `json:"cardUserKey"`
	BinNumber                    string  `json:"binNumber"`
	LastFourDigits               string  `json:"lastFourDigits"`
	BasketId                     string  `json:"basketId"`
	// items
	ConnectorName string `json:"connectorName"`
	AuthCode      string `json:"authCode"`
	HostReference string `json:"hostReference"`
	Phase         string `json:"phase"`
}

func (c CheckoutForm) IsSuccess() bool {
	if c.PaymentStatus == "SUCCESS" {
		return true
	}
	return false
}
