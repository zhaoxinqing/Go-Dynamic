package public

import (
	"bytes"
	"net/http"
)

// Get 发送 GET 请求
func HttpGet(url string) (*http.Response, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 发送 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Post 发送 POST 请求
func HttpPost(url string, body []byte) (*http.Response, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建请求体
	reqBody := bytes.NewBuffer(body)

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	// 发送 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
