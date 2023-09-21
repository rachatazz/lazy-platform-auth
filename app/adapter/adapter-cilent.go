package adapter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type AdapterOptions struct {
	params  map[string]string
	headers map[string]string
}

type Params map[string]string

func HttpGetRequest(
	baseURL string,
	options *AdapterOptions,
) (*http.Response, error) {
	if options.params != nil {
		params := url.Values{}
		for key, value := range options.params {
			params.Add(key, value)
		}
		baseURL = baseURL + "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}

	if options.headers != nil {
		for key, value := range options.headers {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func HttpPostRequestJson(
	baseURL string,
	body *interface{},
	options *AdapterOptions,
) (*http.Response, error) {
	if options.params != nil {
		params := url.Values{}
		for key, value := range options.params {
			params.Add(key, value)
		}
		baseURL = baseURL + "?" + params.Encode()
	}

	jsonStr := []byte(`{}`)
	if body != nil {
		jsonb, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		jsonStr = jsonb
	}

	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if options.headers != nil {
		for key, value := range options.headers {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
