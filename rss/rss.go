package rss

import "GoRss2Webhook/rss/store"

var feedStore *store.FeedStore

func SetFeedStore(store *store.FeedStore) {
	feedStore = store
}

func Subscribe(subscriber store.FeedSubscriber) error {
	s := *feedStore
	return s.Subscribe(subscriber)
}

func Unsubscribe(feedUrl string) error {
	s := *feedStore
	return s.Unsubscribe(feedUrl)
}
