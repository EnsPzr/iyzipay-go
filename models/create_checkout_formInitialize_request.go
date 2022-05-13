package models

import (
	"fmt"
	"github.com/enspzr/iyzipay-go/helpers"
	"strings"
)

type CreateCheckoutFormInitializeRequest struct {
	BaseRequest
	Price           string       `json:"price,omitempty"`
	BasketId        string       `json:"basketId,omitempty"`
	PaymentGroup    string       `json:"paymentGroup,omitempty"`
	Buyer           Buyer        `json:"buyer,omitempty"`
	ShippingAddress Address      `json:"shippingAddress,omitempty"`
	BillingAddress  Address      `json:"billingAddress,omitempty"`
	BasketItems     []BasketItem `json:"basketItems,omitempty"`
	CallbackUrl     string       `json:"callbackUrl,omitempty"`
	Currency        string       `json:"currency,omitempty"`
	PaidPrice       string       `json:"paidPrice,omitempty"`

	PaymentSource       string `json:"paymentSource,omitempty"`
	ForceThreeDS        *int   `json:"forceThreeDS,omitempty"`
	CardUserKey         string `json:"cardUserKey,omitempty"`
	PosOrderId          string `json:"posOrderId,omitempty"`
	EnabledInstallments []int  `json:"enabledInstallments,omitempty"`
}

func (c *CreateCheckoutFormInitializeRequest) String() string {
	t := ""
	t += helpers.Append("locale", c.BaseRequest.Locale)
	t += helpers.Append("conversationId", c.BaseRequest.ConversationId)
	t += helpers.AppendPrice("price", c.Price)
	t += helpers.Append("basketId", c.BasketId)
	t += helpers.Append("paymentGroup", c.PaymentGroup)
	t += helpers.AppendArray("buyer", helpers.GetRequestString(c.Buyer.String()))
	t += helpers.AppendArray("shippingAddress", helpers.GetRequestString(c.ShippingAddress.String()))
	t += helpers.AppendArray("billingAddress", helpers.GetRequestString(c.BillingAddress.String()))
	basket := make([]string, 0)
	for _, item := range c.BasketItems {
		basket = append(basket, fmt.Sprintf("%s", helpers.GetRequestString(item.String())))
	}
	t += helpers.AppendArray("basketItems", helpers.GetRequestString(strings.Join(basket, ", ")+","))
	t += helpers.Append("callbackUrl", c.CallbackUrl)
	t += helpers.Append("currency", c.Currency)
	t += helpers.Append("paymentSource", c.PaymentSource)
	//t += helpers.Append("forceThreeDS", c.ForceThreeDS)
	t += helpers.Append("cardUserKey", c.CardUserKey)
	t += helpers.AppendPrice("paidPrice", c.PaidPrice)
	//t += helpers.AppendArray("enabledInstallments", c.EnabledInstallments)
	return t
}
