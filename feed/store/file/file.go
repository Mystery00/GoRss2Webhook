package file

import (
	"GoRss2Webhook/feed/store"
	"GoRss2Webhook/utils/file"
	"github.com/sirupsen/logrus"
)

type fileStore struct {
	storePath string
	fileName  string
	data      []store.FeedSubscriber
}

func Init(storePath, fileName string) store.FeedStore {
	data := make([]store.FeedSubscriber, 0)
	err := file.Read(storePath, fileName, &data)
	if err != nil {
		logrus.Warnf(`read file error: %s`, err.Error())
	}
	var rssStore store.FeedStore
	fileStore := &fileStore{
		storePath: storePath,
		fileName:  fileName,
		data:      data,
	}
	rssStore = fileStore
	return rssStore
}

func (store *fileStore) Subscribe(subscriber store.FeedSubscriber) error {
	store.data = append(store.data, subscriber)
	return file.Write(store.storePath, store.fileName, store.data)
}

func (store *fileStore) GetAll() ([]store.FeedSubscriber, error) {
	return store.data, nil
}

func (store *fileStore) Unsubscribe(feedUrl string) error {
	for i, subscriber := range store.data {
		if subscriber.FeedUrl == feedUrl {
			store.data = append(store.data[:i], store.data[i+1:]...)
			return nil
		}
	}
	return file.Write(store.storePath, store.fileName, store.data)
}
