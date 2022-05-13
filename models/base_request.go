package models

type BaseRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}
