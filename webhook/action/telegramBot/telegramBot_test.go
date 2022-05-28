package telegramBot

import (
	"GoRss2Webhook/feed/fetch"
	"GoRss2Webhook/feed/store"
	"os"
	"testing"
)

func TestDoWebhook(t *testing.T) {
	metaData := `{"content":"检测到RSS数据更新: {{.Title}} ，访问地址为: {{.Link}}","token":"xxxx","chatId":"xxxx"}`
	subscriber := store.FeedSubscriber{
		FeedUrl: "https://rsshub.admin.mystery0.vip/juejin/category/backend",
	}
	parse, err := fetch.Parse(subscriber)
	if err != nil {
		t.Error(err)
	}

	action := TelegramBotAction{}
	action.DoWebhook(metaData, *parse.Items[0])
	t.Log("send done")
}

func TestDoSend(t *testing.T) {
	w := TelegramBot{
		Content:  `Hello World!`,
		BotToken: os.Getenv("TG_BOT_TOKEN"),
		Host:     os.Getenv("TG_HOST"),
		ChatId:   os.Getenv("TG_CHAT_ID"),
	}
	subscriber := store.FeedSubscriber{
		FeedUrl: "https://rsshub.admin.mystery0.vip/coronavirus/dxy",
	}
	parse, err := fetch.Parse(subscriber)
	if err != nil {
		t.Error(err)
	}

	action := TelegramBotAction{}
	action.DoSend(w, *parse.Items[0])
	t.Log("send done")
}
