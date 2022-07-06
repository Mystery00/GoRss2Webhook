package memory

import (
	"GoRss2Webhook/store"
	"crypto/sha512"
	"encoding/hex"
	"github.com/mmcdole/gofeed"
)

type memoryStore struct {
	data map[string][]string
}

func Init() store.RssStore {
	var rssStore store.RssStore
	rssStore = &memoryStore{
		data: make(map[string][]string),
	}
	return rssStore
}

func hash(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (store *memoryStore) Save(feed gofeed.Feed, item gofeed.Item) error {
	key := hash(feed.Link)
	store.data[key] = append(store.data[key], hash(item.GUID))
	return nil
}

func (store *memoryStore) Exist(feed gofeed.Feed, item gofeed.Item) bool {
	key := hash(feed.Link)
	values := store.data[key]
	if values == nil || len(values) == 0 {
		return false
	}
	var checkValue string
	if item.GUID != "" {
		checkValue = hash(item.GUID)
	} else {
		checkValue = hash(item.Link)
	}
	for _, value := range values {
		if value == checkValue {
			return true
		}
	}
	return false
}

func (store *memoryStore) Delete(feed gofeed.Feed) error {
	key := hash(feed.Link)
	delete(store.data, key)
	return nil
}

func (store *memoryStore) Clear() error {
	store.data = make(map[string][]string)
	return nil
}
