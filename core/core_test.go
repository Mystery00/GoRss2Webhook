package core

import (
	"GoRss2Webhook/rss/store"
	rss "GoRss2Webhook/rss/store/memory"
	feed "GoRss2Webhook/store/memory"
	"GoRss2Webhook/webhook/store/memory"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestDoWork(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	feedStore := rss.Init()
	rssStore := feed.Init()
	webhookStore := memory.Init()

	subscriber := store.FeedSubscriber{
		FeedUrl: "https://rsshub.admin.mystery0.vip/coronavirus/dxy",
	}
	err := feedStore.Subscribe(subscriber)
	if err != nil {
		t.Error(err)
	}
	DoWork(feedStore, rssStore, webhookStore)
	DoWork(feedStore, rssStore, webhookStore)
	time.Sleep(time.Minute)
}
