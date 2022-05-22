package memory

import (
	store3 "GoRss2Webhook/webhook/store"
	"testing"
)

func TestMemory(t *testing.T) {
	store := Init()
	webhook := store3.Webhook{
		SubscribeUrl: "http",
	}
	err := store.Save(webhook)
	if err != nil {
		t.Error(err)
	}
	all, err := store.GetAll("http")
	if err != nil {
		t.Error(err)
	}
	if len(all) != 1 {
		t.Error("Expected 1, got ", len(all))
	}
	err = store.Delete("http")
	if err != nil {
		t.Error(err)
	}
	all, err = store.GetAll("http")
	if err != nil {
		t.Error(err)
	}
	if len(all) != 0 {
		t.Error("Expected 0, got ", len(all))
	}
	webhook = store3.Webhook{
		SubscribeUrl: "https",
	}
	err = store.Save(webhook)
	if err != nil {
		t.Error(err)
	}
	all, err = store.GetAll("https")
	if err != nil {
		t.Error(err)
	}
	if len(all) != 1 {
		t.Error("Expected 1, got ", len(all))
	}
}