package core

import (
	"GoRss2Webhook/feed/store"
	feed "GoRss2Webhook/feed/store/memory"
	rss "GoRss2Webhook/store/memory"
	"GoRss2Webhook/webhook/store/memory"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_doWork(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	feedStore := feed.Init()
	rssStore := rss.Init()
	webhookStore := memory.Init()

	subscriber := store.FeedSubscriber{
		FeedUrl: "https://rsshub.admin.mystery0.vip/coronavirus/dxy",
	}
	err := feedStore.Subscribe(subscriber)
	if err != nil {
		t.Error(err)
	}
	doWork(feedStore, rssStore, webhookStore)
	doWork(feedStore, rssStore, webhookStore)
	time.Sleep(time.Minute)
}
