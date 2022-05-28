package conf

import (
	store3 "GoRss2Webhook/webhook/store"
	"os"
	"testing"
)

func TestYamlFile(t *testing.T) {
	storePath := `/tmp/GoRss2Webhook/webhook_store`
	store := Init(storePath, `conf`, `yaml`)
	webhook := store3.Webhook{}
	err := store.Save("http", webhook)
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
	webhook = store3.Webhook{}
	err = store.Save("https", webhook)
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
	err = os.RemoveAll(storePath)
	if err != nil {
		t.Error(err)
	}
}

func TestJsonFile(t *testing.T) {
	storePath := `/tmp/GoRss2Webhook/webhook_store`
	store := Init(storePath, `conf`, `json`)
	webhook := store3.Webhook{}
	err := store.Save("http", webhook)
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
	webhook = store3.Webhook{}
	err = store.Save("https", webhook)
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
	err = os.RemoveAll(storePath)
	if err != nil {
		t.Error(err)
	}
}
