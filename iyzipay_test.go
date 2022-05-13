package iyzipay

import (
	"github.com/enspzr/iyzipay-go/helpers"
	"github.com/enspzr/iyzipay-go/models"
	"os"
	"testing"
	"time"
)

func TestIyzipaySingleBasketItem(t *testing.T) {
	options := models.Options{
		ApiKey:    os.Getenv("API_KEY"),
		SecretKey: os.Getenv("SECRET_KEY"),
		BaseUrl:   "https://sandbox-api.iyzipay.com",
	}

	time := time.Date(2001, 9, 21, 12, 0, 0, 0, time.Local)
	request := models.CreateCheckoutFormInitializeRequest{
		BaseRequest: models.BaseRequest{
			Locale:         "tr",
			ConversationId: "123456",
		},
		Price:        "450.0",
		PaidPrice:    "450.150",
		BasketId:     "B123sr",
		PaymentGroup: "PRODUCT",
		Currency:     "TRY",
		Buyer: models.Buyer{
			Id:                  "10",
			Name:                "Enes",
			Surname:             "PAZAR",
			IdentityNumber:      "11111111111",
			Email:               "m.enespazar@gmail.com",
			GsmNumber:           "05355555555",
			RegistrationDate:    helpers.IyzipayTimeString(time),
			LastLoginDate:       helpers.IyzipayTimeString(time),
			RegistrationAddress: "Kutahya",
			City:                "Kutahya",
			Country:             "Turkiye",
			ZipCode:             "43000",
			Ip:                  "127.0.0.1",
		},
		ShippingAddress: models.Address{
			Description: "Kutahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kutahya",
			Country:     "Turkiye",
		},
		BillingAddress: models.Address{
			Description: "Kutahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kutahya",
			Country:     "Turkiye",
		},
		BasketItems: make([]models.BasketItem, 0),
		CallbackUrl: "http://localhost:8080",
	}

	request.BasketItems = append(request.BasketItems, models.BasketItem{
		Id:               "101",
		Price:            "450",
		Name:             "Tursu",
		Category1:        "Software",
		Category2:        "Electronics",
		ItemType:         "VIRTUAL",
		SubMerchantKey:   "",
		SubMerchantPrice: "",
	})

	body, _ := CheckoutFormCreate(request, options)
	if body.Status != "success" {
		t.Error("tekli sepet success değil")
	}
}

func TestIyzipayMultiBasketItem(t *testing.T) {
	options := models.Options{
		ApiKey:    os.Getenv("API_KEY"),
		SecretKey: os.Getenv("SECRET_KEY"),
		BaseUrl:   "https://sandbox-api.iyzipay.com",
	}

	time := time.Date(2001, 9, 21, 12, 0, 0, 0, time.Local)
	request := models.CreateCheckoutFormInitializeRequest{
		BaseRequest: models.BaseRequest{
			Locale:         "tr",
			ConversationId: "123456",
		},
		Price:        "900.0",
		PaidPrice:    "1050.150",
		BasketId:     "B123sr",
		PaymentGroup: "PRODUCT",
		Currency:     "TRY",
		Buyer: models.Buyer{
			Id:                  "10",
			Name:                "Enes",
			Surname:             "PAZAR",
			IdentityNumber:      "11111111111",
			Email:               "m.enespazar@gmail.com",
			GsmNumber:           "05355555555",
			RegistrationDate:    helpers.IyzipayTimeString(time),
			LastLoginDate:       helpers.IyzipayTimeString(time),
			RegistrationAddress: "Kutahya",
			City:                "Kutahya",
			Country:             "Turkiye",
			ZipCode:             "43000",
			Ip:                  "127.0.0.1",
		},
		ShippingAddress: models.Address{
			Description: "Kutahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kutahya",
			Country:     "Turkiye",
		},
		BillingAddress: models.Address{
			Description: "Kutahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kutahya",
			Country:     "Turkiye",
		},
		BasketItems: make([]models.BasketItem, 0),
		CallbackUrl: "http://localhost:8080",
	}

	request.BasketItems = append(request.BasketItems, models.BasketItem{
		Id:               "101",
		Price:            "450",
		Name:             "Tursu",
		Category1:        "Software",
		Category2:        "Electronics",
		ItemType:         "VIRTUAL",
		SubMerchantKey:   "",
		SubMerchantPrice: "",
	})

	request.BasketItems = append(request.BasketItems, models.BasketItem{
		Id:               "101",
		Price:            "450",
		Name:             "Lahmacun",
		Category1:        "Software",
		Category2:        "Electronics",
		ItemType:         "VIRTUAL",
		SubMerchantKey:   "",
		SubMerchantPrice: "",
	})

	body, _ := CheckoutFormCreate(request, options)
	if body.Status != "success" {
		t.Error("çoklu sepet success değil")
	}
}

func TestIyzipayTurkceKarakter(t *testing.T) {
	options := models.Options{
		ApiKey:    os.Getenv("API_KEY"),
		SecretKey: os.Getenv("SECRET_KEY"),
		BaseUrl:   "https://sandbox-api.iyzipay.com",
	}

	time := time.Date(2001, 9, 21, 12, 0, 0, 0, time.Local)
	request := models.CreateCheckoutFormInitializeRequest{
		BaseRequest: models.BaseRequest{
			Locale:         "tr",
			ConversationId: "123456",
		},
		Price:        "900.0",
		PaidPrice:    "1050.150",
		BasketId:     "B123sr",
		PaymentGroup: "PRODUCT",
		Currency:     "TRY",
		Buyer: models.Buyer{
			Id:                  "10",
			Name:                "Enes",
			Surname:             "PAZAR",
			IdentityNumber:      "11111111111",
			Email:               "m.enespazar@gmail.com",
			GsmNumber:           "05355555555",
			RegistrationDate:    helpers.IyzipayTimeString(time),
			LastLoginDate:       helpers.IyzipayTimeString(time),
			RegistrationAddress: "Kütahya",
			City:                "Kütahya",
			Country:             "Türkiye",
			ZipCode:             "43000",
			Ip:                  "127.0.0.1",
		},
		ShippingAddress: models.Address{
			Description: "Kütahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kütahya",
			Country:     "Türkiye",
		},
		BillingAddress: models.Address{
			Description: "Kütahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kütahya",
			Country:     "Türkiye",
		},
		BasketItems: make([]models.BasketItem, 0),
		CallbackUrl: "http://localhost:8080",
	}

	request.BasketItems = append(request.BasketItems, models.BasketItem{
		Id:               "101",
		Price:            "450",
		Name:             "Turşu",
		Category1:        "Software",
		Category2:        "Electronics",
		ItemType:         "VIRTUAL",
		SubMerchantKey:   "",
		SubMerchantPrice: "",
	})

	request.BasketItems = append(request.BasketItems, models.BasketItem{
		Id:               "101",
		Price:            "450",
		Name:             "Lahmacun",
		Category1:        "Software",
		Category2:        "Electronics",
		ItemType:         "VIRTUAL",
		SubMerchantKey:   "",
		SubMerchantPrice: "",
	})

	body, _ := CheckoutFormCreate(request, options)
	if body.Status != "success" {
		t.Error("çoklu sepet success değil")
	}
}

func TestIyzipayTurkceKarakterNotConversationId(t *testing.T) {
	options := models.Options{
		ApiKey:    os.Getenv("API_KEY"),
		SecretKey: os.Getenv("SECRET_KEY"),
		BaseUrl:   "https://sandbox-api.iyzipay.com",
	}

	time := time.Date(2001, 9, 21, 12, 0, 0, 0, time.Local)
	request := models.CreateCheckoutFormInitializeRequest{
		BaseRequest: models.BaseRequest{
			Locale: "tr",
		},
		Price:        "900.0",
		PaidPrice:    "1050.150",
		BasketId:     "B123sr",
		PaymentGroup: "PRODUCT",
		Currency:     "TRY",
		Buyer: models.Buyer{
			Id:                  "10",
			Name:                "Enes",
			Surname:             "PAZAR",
			IdentityNumber:      "11111111111",
			Email:               "m.enespazar@gmail.com",
			GsmNumber:           "05355555555",
			RegistrationDate:    helpers.IyzipayTimeString(time),
			LastLoginDate:       helpers.IyzipayTimeString(time),
			RegistrationAddress: "Kütahya",
			City:                "Kütahya",
			Country:             "Türkiye",
			ZipCode:             "43000",
			Ip:                  "127.0.0.1",
		},
		ShippingAddress: models.Address{
			Description: "Kütahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kütahya",
			Country:     "Türkiye",
		},
		BillingAddress: models.Address{
			Description: "Kütahya Merkez",
			ZipCode:     "43000",
			ContactName: "Enes PAZAR",
			City:        "Kütahya",
			Country:     "Türkiye",
		},
		BasketItems: make([]models.BasketItem, 0),
		CallbackUrl: "http://localhost:8080",
	}

	request.BasketItems = append(request.BasketItems, models.BasketItem{
		Id:               "101",
		Price:            "450",
		Name:             "Turşu",
		Category1:        "Software",
		Category2:        "Electronics",
		ItemType:         "VIRTUAL",
		SubMerchantKey:   "",
		SubMerchantPrice: "",
	})

	request.BasketItems = append(request.BasketItems, models.BasketItem{
		Id:               "101",
		Price:            "450",
		Name:             "Lahmacun",
		Category1:        "Software",
		Category2:        "Electronics",
		ItemType:         "VIRTUAL",
		SubMerchantKey:   "",
		SubMerchantPrice: "",
	})

	body, _ := CheckoutFormCreate(request, options)
	if body.Status != "success" {
		t.Error("çoklu sepet success değil")
	}
}
