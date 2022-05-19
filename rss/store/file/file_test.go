package file

import (
	store2 "GoRss2Webhook/rss/store"
	"testing"
)

func TestFile(t *testing.T) {
	store := Init(`/tmp/GoRss2Webhook/rss`, `rss.json`, "*/10 * * * * ?")
	subscriber := store2.FeedSubscriber{
		FeedUrl:   "http",
		UserAgent: "",
		ProxyUrl:  "",
	}
	err := store.Save(subscriber)
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
	err = store.Delete("http")
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
	err = store.Save(subscriber)
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
