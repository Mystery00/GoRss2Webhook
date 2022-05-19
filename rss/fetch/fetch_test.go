package fetch

import (
	"GoRss2Webhook/rss/store"
	"testing"
)

func TestParse(t *testing.T) {
	subscriber := store.FeedSubscriber{
		FeedUrl: "https://rsshub.admin.mystery0.vip/coronavirus/dxy",
	}
	parse, err := Parse(subscriber)
	if err != nil {
		t.Error(err)
	}
	t.Log(parse)
}
