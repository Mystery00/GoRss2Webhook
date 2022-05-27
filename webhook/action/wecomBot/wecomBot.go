package wecomBot

import (
	"GoRss2Webhook/webhook/action/http"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/tidwall/gjson"
)

type WecomBot struct {
	Content string
	Body    string
	Key     string
	Url     string
}

type WecomBotAction struct {
}

func (a WecomBotAction) DoWebhook(metaData string, item gofeed.Item) {
	w := WecomBot{
		Content: gjson.Get(metaData, "content").String(),
		Body:    gjson.Get(metaData, "body").String(),
		Key:     gjson.Get(metaData, "key").String(),
		Url:     gjson.Get(metaData, "url").String(),
	}
	a.DoSend(w, item)
}

func (a WecomBotAction) DoSend(w WecomBot, item gofeed.Item) {
	u := fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s`, w.Key)
	if w.Url != "" {
		u = w.Url
	}
	header := make(map[string]string, 1)
	header["Content-Type"] = "application/json"
	content := fmt.Sprintf(`{"msgtype": "text","text": {"content":"%s"}}`, w.Content)
	if w.Body != "" {
		content = w.Body
	}
	h := http.Http{
		Url:    u,
		Method: "POST",
		Body:   content,
		Header: header,
	}
	httpAction := http.HttpAction{}
	httpAction.DoSend(h, item)
}
