package http

import (
	"GoRss2Webhook/webhook/action/tool"
	"bytes"
	"encoding/json"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Http struct {
	Url      string
	Method   string
	Body     string
	Header   map[string]string
	ProxyUrl string
	Timeout  time.Duration
}

type HttpAction struct {
}

func (a HttpAction) DoWebhook(metaData string, item gofeed.Item) {
	header := make(map[string]string, 0)
	err := json.Unmarshal([]byte(gjson.Get(metaData, ".header").String()), &header)
	if err != nil {
		logrus.Warnf(`parse http header failed, %s`, err)
	}
	duration, err := time.ParseDuration(gjson.Get(metaData, ".url").String())
	if err != nil {
		duration = 0
	}
	h := Http{
		Url:      gjson.Get(metaData, ".url").String(),
		Method:   gjson.Get(metaData, ".method").Str,
		Body:     gjson.Get(metaData, ".body").String(),
		Header:   header,
		ProxyUrl: gjson.Get(metaData, ".proxyUrl").String(),
		Timeout:  duration,
	}
	a.DoSend(h, item)
}

func (a HttpAction) DoSend(h Http, item gofeed.Item) {
	body := tool.ParseTemplate(h.Body, item)
	urlString := h.Url
	method := strings.ToUpper(h.Method)
	headers := h.Header

	proxy := http.ProxyFromEnvironment
	if h.ProxyUrl != "" {
		p, _ := url.Parse(h.ProxyUrl)
		proxy = http.ProxyURL(p)
	}
	timeout := h.Timeout
	if timeout <= 0 {
		timeout = time.Second * 5
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
		},
		Timeout: timeout,
	}
	sendHttp(client, urlString, method, body, headers)
}

func sendHttp(client *http.Client, url, method string, body []byte, headers map[string]string) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	//设置header参数
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//发送请求
	logrus.Debugf(`send http request, %s %s, body: %s, header: %s`, method, url, string(body), headers)
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
