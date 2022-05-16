package memory

import (
	"GoRss2Webhook/rss/store"
)

type memoryStore struct {
	data []string
}

func Init() store.FeedStore {
	var rssStore store.FeedStore
	rssStore = &memoryStore{
		data: make([]string, 0),
	}
	return rssStore
}

func (store *memoryStore) Save(feedUrl string) error {
	store.data = append(store.data, feedUrl)
	return nil
}

func (store *memoryStore) GetAll() ([]string, error) {
	return store.data, nil
}

func (store *memoryStore) Delete(feedUrl string) error {
	for i, url := range store.data {
		if url == feedUrl {
			store.data = append(store.data[:i], store.data[i+1:]...)
			return nil
		}
	}
	return nil
}
