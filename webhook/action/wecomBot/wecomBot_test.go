package wecomBot

import (
	"GoRss2Webhook/feed/fetch"
	"GoRss2Webhook/feed/store"
	"os"
	"testing"
)

func TestDoSend(t *testing.T) {
	w := WecomBot{
		Content: `{"msgtype": "text","text": {"content":"Hello World!"}}`,
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
