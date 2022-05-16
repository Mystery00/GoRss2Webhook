package file

import (
	"github.com/mmcdole/gofeed"
	"testing"
)

func TestFile(t *testing.T) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://rsshub.admin.mystery0.vip/coronavirus/dxy")
	store := Init(`/tmp/GoRss2Webhook/rss_store`)
	exist := store.Exist(*feed, *feed.Items[0])
	if exist {
		t.Error("Why Exist ?!")
	}
	err := store.Save(*feed, *feed.Items[0])
	if err != nil {
		t.Error(err)
	}
	exist = store.Exist(*feed, *feed.Items[0])
	if !exist {
		t.Error("Why not Exist ?!")
	}
	exist = store.Exist(*feed, *feed.Items[1])
	if exist {
		t.Error("Why Exist ?!")
	}
	err = store.Save(*feed, *feed.Items[1])
	if err != nil {
		t.Error(err)
	}
	err = store.Delete(*feed)
	if err != nil {
		t.Error(err)
	}
	err = store.Clear()
	if err != nil {
		t.Error(err)
	}
	exist = store.Exist(*feed, *feed.Items[0])
	if exist {
		t.Error("Why Exist ?!")
	}
	exist = store.Exist(*feed, *feed.Items[1])
	if exist {
		t.Error("Why Exist ?!")
	}
}
