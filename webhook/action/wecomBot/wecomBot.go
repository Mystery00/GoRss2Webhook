package wecomBot

import (
	"GoRss2Webhook/webhook/action/http"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/tidwall/gjson"
)

type WecomBot struct {
	Content string
	Key     string
	Url     string
}

type WecomBotAction struct {
}

func (a WecomBotAction) DoWebhook(metaData string, item gofeed.Item) {
	w := WecomBot{
		Content: gjson.Get(metaData, ".content").String(),
		Key:     gjson.Get(metaData, ".key").String(),
		Url:     gjson.Get(metaData, ".url").String(),
	}
	a.DoSend(w, item)
}

func (a WecomBotAction) DoSend(w WecomBot, item gofeed.Item) {
	u := w.Url
	if u == "" {
		u = fmt.Sprintf(`https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s`, w.Key)
	}
	header := make(map[string]string, 1)
	header["Content-Type"] = "application/json"
	h := http.Http{
		Url:    u,
		Method: "POST",
		Body:   w.Content,
		Header: header,
	}
	httpAction := http.HttpAction{}
	httpAction.DoSend(h, item)
}
