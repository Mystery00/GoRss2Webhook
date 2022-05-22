package memory

import (
	store2 "GoRss2Webhook/feed/store"
	"testing"
)

func TestMemory(t *testing.T) {
	store := Init()
	subscriber := store2.FeedSubscriber{
		FeedUrl:   "http",
		UserAgent: "",
		ProxyUrl:  "",
	}
	err := store.Subscribe(subscriber)
	if err != nil {
		t.Error(err)
	}
	all, err := store.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(all) != 1 {
		t.Error("Expected 1, got ", len(all))
	}
	if all[0].FeedUrl != "http" {
		t.Error("Expected http, got ", all[0])
	}
	err = store.Unsubscribe("http")
	if err != nil {
		t.Error(err)
	}
	all, err = store.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(all) != 0 {
		t.Error("Expected 0, got ", len(all))
	}
	subscriber = store2.FeedSubscriber{
		FeedUrl:   "https",
		UserAgent: "",
		ProxyUrl:  "",
	}
	err = store.Subscribe(subscriber)
	if err != nil {
		t.Error(err)
	}
	all, err = store.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(all) != 1 {
		t.Error("Expected 1, got ", len(all))
	}
	if all[0].FeedUrl != "https" {
		t.Error("Expected https, got ", all[0])
	}
}
