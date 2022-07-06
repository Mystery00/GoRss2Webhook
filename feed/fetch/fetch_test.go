package fetch

import (
	"GoRss2Webhook/feed/store"
	"testing"
)

func TestParse(t *testing.T) {
	subscriber := store.FeedSubscriber{
		FeedUrl: "https://sspai.com/feed",
	}
	parse, err := Parse(subscriber)
	if err != nil {
		t.Error(err)
	}
	t.Log(parse)
}
