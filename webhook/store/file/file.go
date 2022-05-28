package file

import (
	"GoRss2Webhook/utils/file"
	"GoRss2Webhook/webhook/store"
	"github.com/sirupsen/logrus"
)

type fileStore struct {
	storePath string
	fileName  string
	data      map[string][]store.Webhook
}

func Init(storePath, fileName string) store.WebhookStore {
	data := make(map[string][]store.Webhook, 0)
	err := file.Read(storePath, fileName, &data)
	if err != nil {
		logrus.Warnf(`read file error: %s`, err.Error())
	}
	var webhookStore store.WebhookStore
	fileStore := &fileStore{
		storePath: storePath,
		fileName:  fileName,
		data:      data,
	}
	webhookStore = fileStore
	return webhookStore
}

func (store *fileStore) Save(feedUrl string, webhook store.Webhook) error {
	key := file.Hash(feedUrl)
	store.data[key] = append(store.data[key], webhook)
	return file.Write(store.storePath, store.fileName, store.data)
}

func (store *fileStore) GetAll(feedUrl string) ([]store.Webhook, error) {
	key := file.Hash(feedUrl)
	return store.data[key], nil
}

func (store *fileStore) Delete(feedUrl string) error {
	key := file.Hash(feedUrl)
	delete(store.data, key)
	return file.Write(store.storePath, store.fileName, store.data)
}
