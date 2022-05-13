package models

import "github.com/enspzr/iyzipay-go/helpers"

type RetrieveCheckoutFormRequest struct {
	BaseRequest
	Token string `json:"token"`
}

func (r RetrieveCheckoutFormRequest) String() string {
	t := ""
	t += helpers.Append("locale", r.BaseRequest.Locale)
	t += helpers.Append("conversationId", r.BaseRequest.ConversationId)
	t += helpers.Append("token", r.Token)
	return t
}
