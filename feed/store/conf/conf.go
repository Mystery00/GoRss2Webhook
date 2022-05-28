package conf

import (
	"GoRss2Webhook/feed/store"
	"GoRss2Webhook/utils/conf"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type confStore struct {
	storePath  string
	fileName   string
	configType string
	data       subscriberList
}

type subscriberList struct {
	Subscriber []store.FeedSubscriber `mapstructure:"subscriber"`
}

func Init(storePath, fileName, configType string) store.FeedStore {
	var list subscriberList
	c, err := conf.UnmarshalFromFile(&list, storePath, fileName, configType)
	if err != nil {
		logrus.Warnf(`parse conf error: %s`, err.Error())
	}
	var feedStore store.FeedStore
	configStore := &confStore{
		storePath:  storePath,
		fileName:   fileName,
		configType: configType,
		data:       list,
	}
	conf.WatchChange(c, func(v *viper.Viper) {
		var list subscriberList
		_, err := conf.UnmarshalFromFile(&list, storePath, fileName, configType)
		if err != nil {
			logrus.Warnf(`parse conf error: %s`, err.Error())
		}
		configStore.data = list
	})
	feedStore = configStore
	return feedStore
}

func (store *confStore) Subscribe(subscriber store.FeedSubscriber) error {
	for i, subscriber := range store.data.Subscriber {
		if subscriber.FeedUrl == subscriber.FeedUrl {
			store.data.Subscriber = append(store.data.Subscriber[:i], store.data.Subscriber[i+1:]...)
			break
		}
	}
	store.data.Subscriber = append(store.data.Subscriber, subscriber)
	return conf.MarshalToFile(store.data, store.storePath, store.fileName, store.configType)
}

func (store *confStore) GetAll() ([]store.FeedSubscriber, error) {
	return store.data.Subscriber, nil
}

func (store *confStore) Unsubscribe(feedUrl string) error {
	for i, subscriber := range store.data.Subscriber {
		if subscriber.FeedUrl == feedUrl {
			store.data.Subscriber = append(store.data.Subscriber[:i], store.data.Subscriber[i+1:]...)
			return nil
		}
	}
	return conf.MarshalToFile(store.data, store.storePath, store.fileName, store.configType)
}
