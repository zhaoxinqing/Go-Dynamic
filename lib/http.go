package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Get http get method
func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	//http client
	client := &http.Client{}
	return client.Do(req)
}

// Post http post method
func Post2(url string, body map[string]string, params map[string]string, headers map[string]string) (*http.Response, error) {
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	//http client
	client := &http.Client{}
	return client.Do(req)
}
