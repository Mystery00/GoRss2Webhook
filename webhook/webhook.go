package webhook

import (
	"GoRss2Webhook/webhook/store"
	"bytes"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"
)

func DoWebhook(webhook store.Webhook, item gofeed.Item) {
	body := parseTemplate(webhook.Http.Body, item)
	urlString := webhook.Http.Url
	method := strings.ToUpper(webhook.Http.Method)
	headers := webhook.Http.Header

	proxy := http.ProxyFromEnvironment
	if webhook.ProxyUrl != "" {
		p, _ := url.Parse(webhook.ProxyUrl)
		proxy = http.ProxyURL(p)
	}
	timeout := webhook.Timeout
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

func parseTemplate(text string, item gofeed.Item) []byte {
	// 根据指定模版文本生成handler
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		panic(err)
	}
	// 模版渲染，并赋值给变量
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, item); err != nil {
		panic(err)
	}
	return buf.Bytes()
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
