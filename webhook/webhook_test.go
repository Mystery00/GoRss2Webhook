package webhook

import (
	"github.com/mmcdole/gofeed"
	"net/http"
	"testing"
)

func Test_parseTemplate(t *testing.T) {
	tpl := `标题: {{.Title}} 描述: {{.Description}}`
	item := gofeed.Item{
		Title:       "国家卫健委",
		Description: "5月21日0—24时",
	}
	bytes := parseTemplate(tpl, item)
	if string(bytes) != "标题: 国家卫健委 描述: 5月21日0—24时" {
		t.Error(`parseTemplate failed`)
	}
}

func Test_sendHttp(t *testing.T) {
	url := "https://www.baidu.com"
	method := "POST"
	body := []byte("hhhhh")
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	sendHttp(&http.Client{}, url, method, body, headers)
}
