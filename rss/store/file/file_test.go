package file

import "testing"

func TestFile(t *testing.T) {
	store := Init(`/tmp/GoRss2Webhook/rss`, `rss.json`, "*/10 * * * * ?")
	err := store.Save("http")
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
	if all[0] != "http" {
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
	err = store.Save("https")
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
	if all[0] != "https" {
		t.Error("Expected https, got ", all[0])
	}
}
