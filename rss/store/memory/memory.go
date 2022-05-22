package memory

import (
	"GoRss2Webhook/rss/store"
)

type memoryStore struct {
	data []store.FeedSubscriber
}

func Init() store.FeedStore {
	var rssStore store.FeedStore
	rssStore = &memoryStore{
		data: make([]store.FeedSubscriber, 0),
	}
	return rssStore
}

func (store *memoryStore) Subscribe(subscriber store.FeedSubscriber) error {
	store.data = append(store.data, subscriber)
	return nil
}

func (store *memoryStore) GetAll() ([]store.FeedSubscriber, error) {
	return store.data, nil
}

func (store *memoryStore) Unsubscribe(feedUrl string) error {
	for i, subscriber := range store.data {
		if subscriber.FeedUrl == feedUrl {
			store.data = append(store.data[:i], store.data[i+1:]...)
			return nil
		}
	}
	return nil
}
