package iyzipay

import (
	"encoding/json"
	"github.com/enspzr/iyzipay-go/models"
	"io"
)

func CheckoutFormCreate(form models.CreateCheckoutFormInitializeRequest, options models.Options) (models.CheckoutFormInitialize, error) {
	headers := models.GetHttpHeaders(form, options)
	r, w := io.Pipe()
	go func() {
		json.NewEncoder(w).Encode(form)
		w.Close()
	}()
	result, err := post(options.BaseUrl+"/payment/iyzipos/checkoutform/initialize/auth/ecom", &headers, r)
	if err != nil {
		return models.CheckoutFormInitialize{}, err
	}
	formInitialize := models.CheckoutFormInitialize{}
	err = json.Unmarshal(result, &formInitialize)
	if err != nil {
		return formInitialize, err
	}
	return formInitialize, nil
}

func CheckoutFormRetrieve(form models.RetrieveCheckoutFormRequest, options models.Options) (models.CheckoutForm, error) {
	headers := models.GetHttpHeadersRetrieve(form, options)
	r, w := io.Pipe()
	go func() {
		json.NewEncoder(w).Encode(form)
		w.Close()
	}()
	result, err := post(options.BaseUrl+"/payment/iyzipos/checkoutform/auth/ecom/detail", &headers, r)
	if err != nil {
		return models.CheckoutForm{}, err
	}
	checkoutForm := models.CheckoutForm{}
	err = json.Unmarshal(result, &checkoutForm)
	if err != nil {
		return checkoutForm, err
	}
	return checkoutForm, nil
}
