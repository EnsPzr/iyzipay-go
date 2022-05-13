package models

import (
	"fmt"
	"github.com/enspzr/iyzipay-go/helpers"
	"time"
)

type IyzipayResource struct {
	Authorization           string `json:"authorization" default:"Authorization"`
	RandomHeaderName        string `json:"randomHeaderName" default:"x-iyzi-rnd"`
	ClientVersionHeaderName string `json:"clientVersionHeaderName" default:"x-iyzi-client-version"`
	IyziwsHeaderName        string `json:"iyziwsHeaderName" default:"IYZWS"`
	Colon                   string `json:"colon" default:":"`
	Status                  string `json:"status"`
	ErrorCode               string `json:"errorCode"`
	ErrorMessage            string `json:"errorMessage"`
	ErrorGroup              string `json:"errorGroup"`
	Locale                  string `json:"locale"`
	SystemTime              int64  `json:"systemTime"`
	ConversationId          string `json:"conversationId"`
}

func IyzipayResourceInit() IyzipayResource {
	return IyzipayResource{
		Authorization:           "Authorization",
		RandomHeaderName:        "x-iyzi-rnd",
		ClientVersionHeaderName: "x-iyzi-client-version",
		IyziwsHeaderName:        "IYZWS ",
		Colon:                   ":",
	}
}

func GetHttpHeaders(form CreateCheckoutFormInitializeRequest, options Options) map[string]string {
	iyzicoResource := IyzipayResourceInit()
	randomString := getRandomString()
	m := make(map[string]string)
	m["Accept"] = "application/json"
	m[iyzicoResource.RandomHeaderName] = randomString
	m[iyzicoResource.ClientVersionHeaderName] = "iyzipay-dotnet-2.1.39"
	m[iyzicoResource.Authorization] = PrepareAuthorizationString(form, randomString, options)
	m["Content-Type"] = "application/json; charset=utf-16"
	return m
}

func GetHttpHeadersRetrieve(form RetrieveCheckoutFormRequest, options Options) map[string]string {
	iyzicoResource := IyzipayResourceInit()
	randomString := getRandomString()
	m := make(map[string]string)
	m["Accept"] = "application/json"
	m[iyzicoResource.RandomHeaderName] = randomString
	m[iyzicoResource.ClientVersionHeaderName] = "iyzipay-dotnet-2.1.39"
	m[iyzicoResource.Authorization] = PrepareAuthorizationStringRetrieve(form, randomString, options)
	m["Content-Type"] = "application/json; charset=utf-16"
	return m
}

func PrepareAuthorizationString(form CreateCheckoutFormInitializeRequest, randomString string, options Options) string {
	iyzicoResource := IyzipayResourceInit()
	hash := helpers.GenerateHash(options.ApiKey, options.SecretKey, randomString, helpers.GetRequestString(form.String()))
	return fmt.Sprintf("%s%s%s%s", iyzicoResource.IyziwsHeaderName, options.ApiKey, iyzicoResource.Colon, hash)
}

func PrepareAuthorizationStringRetrieve(form RetrieveCheckoutFormRequest, randomString string, options Options) string {
	iyzicoResource := IyzipayResourceInit()
	hash := helpers.GenerateHash(options.ApiKey, options.SecretKey, randomString, helpers.GetRequestString(form.String()))
	return fmt.Sprintf("%s%s%s%s", iyzicoResource.IyziwsHeaderName, options.ApiKey, iyzicoResource.Colon, hash)
}

func getRandomString() string {
	now := time.Now()
	_ = now
	dd := now.Format("02012006150405")
	return dd
}
