package iyzipay

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func get(url string, headers *map[string]string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("İstek sırasında hata context => %s", err.Error()))
	}

	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req.WithContext(ctx))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("İstek sırasında hata response => %s", err.Error()))
	}
	defer resp.Body.Close()
	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	return data, err
}

func post(url string, headers *map[string]string, body *io.PipeReader) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("İstek sırasında hata context => %s", err.Error()))
	}
	if headers != nil {
		for key, value := range *headers {
			req.Header[key] = []string{value}
		}
	}
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req.WithContext(ctx))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("İstek sırasında hata response => %s", err.Error()))
	}
	defer resp.Body.Close()
	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	return data, err
}
