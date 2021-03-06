package wecomBot

import (
	"GoRss2Webhook/feed/fetch"
	"GoRss2Webhook/feed/store"
	"os"
	"testing"
)

func TestDoWebhook(t *testing.T) {
	metaData := `{"content":"检测到RSS数据更新: {{.Title}} ，访问地址为: {{.Link}}","key":"xxx"}`
	subscriber := store.FeedSubscriber{
		FeedUrl: "https://rsshub.admin.mystery0.vip/juejin/category/backend",
	}
	parse, err := fetch.Parse(subscriber)
	if err != nil {
		t.Error(err)
	}

	action := WecomBotAction{}
	action.DoWebhook(metaData, *parse.Items[0])
	t.Log("send done")
}

func TestDoSend(t *testing.T) {
	w := WecomBot{
		Content: `Hello World!`,
		Key:     os.Getenv("QW_KEY"),
	}
	subscriber := store.FeedSubscriber{
		FeedUrl: "https://rsshub.admin.mystery0.vip/coronavirus/dxy",
	}
	parse, err := fetch.Parse(subscriber)
	if err != nil {
		t.Error(err)
	}

	action := WecomBotAction{}
	action.DoSend(w, *parse.Items[0])
	t.Log("send done")
}
