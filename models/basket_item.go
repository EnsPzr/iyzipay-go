package models

import "github.com/enspzr/iyzipay-go/helpers"

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

func (b *BasketItem) String() string {
	t := ""
	t += helpers.Append("id", b.Id)
	t += helpers.AppendPrice("price", b.Price)
	t += helpers.Append("name", b.Name)
	t += helpers.Append("category1", b.Category1)
	t += helpers.Append("category2", b.Category2)
	t += helpers.Append("itemType", b.ItemType)
	t += helpers.Append("subMerchantKey", b.SubMerchantKey)
	t += helpers.AppendPrice("subMerchantPrice", b.SubMerchantKey)
	return t
}
