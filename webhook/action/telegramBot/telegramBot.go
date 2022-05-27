package wecomBot

import (
	"GoRss2Webhook/webhook/action/http"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/tidwall/gjson"
)

type TelegramBot struct {
	Content  string
	Body     string
	BotToken string
	Host     string
	ChatId   string
}

type TelegramBotAction struct {
}

func (a TelegramBotAction) DoWebhook(metaData string, item gofeed.Item) {
	w := TelegramBot{
		Content:  gjson.Get(metaData, "content").String(),
		Body:     gjson.Get(metaData, "body").String(),
		BotToken: gjson.Get(metaData, "token").String(),
		Host:     gjson.Get(metaData, "host").String(),
		ChatId:   gjson.Get(metaData, "chatId").String(),
	}
	a.DoSend(w, item)
}

func (a TelegramBotAction) DoSend(w TelegramBot, item gofeed.Item) {
	host := `https://api.telegram.org`
	if w.Host != "" {
		host = `https://api.telegram.org`
	}
	u := fmt.Sprintf(`%s/bot%s/sendMessage`, host, w.BotToken)
	header := make(map[string]string, 1)
	header["Content-Type"] = "application/json"
	content := fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, w.ChatId, w.Content)
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
