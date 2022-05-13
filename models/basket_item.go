package models

type BasketItem struct {
	Id               string `json:"id,omitempty"`
	Price            string `json:"price,omitempty"`
	Name             string `json:"name,omitempty"`
	Category1        string `json:"category1,omitempty"`
	Category2        string `json:"category2,omitempty"`
	ItemType         string `json:"itemType,omitempty"`
	SubMerchantKey   string `json:"subMerchantKey,omitempty"`
	SubMerchantPrice string `json:"subMerchantPrice,omitempty"`
}
