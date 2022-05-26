package http

import (
	"net/http"
	"testing"
)

func Test_sendHttp(t *testing.T) {
	url := "https://www.baidu.com"
	method := "POST"
	body := []byte("hhhhh")
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	sendHttp(&http.Client{}, url, method, body, headers)
}
