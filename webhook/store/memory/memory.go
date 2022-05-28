package memory

import (
	"GoRss2Webhook/utils/file"
	"GoRss2Webhook/webhook/store"
)

type memoryStore struct {
	data map[string][]store.Webhook
}

func Init() store.WebhookStore {
	var webhookStore store.WebhookStore
	webhookStore = &memoryStore{
		data: make(map[string][]store.Webhook, 0),
	}
	return webhookStore
}

func (store *memoryStore) Save(feedUrl string, webhook store.Webhook) error {
	key := file.Hash(feedUrl)
	store.data[key] = append(store.data[key], webhook)
	return nil
}

func (store *memoryStore) GetAll(feedUrl string) ([]store.Webhook, error) {
	key := file.Hash(feedUrl)
	return store.data[key], nil
}

func (store *memoryStore) Delete(feedUrl string) error {
	key := file.Hash(feedUrl)
	delete(store.data, key)
	return nil
}
