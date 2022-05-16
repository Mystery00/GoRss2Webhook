package rss

import "GoRss2Webhook/rss/store"

var feedStore *store.FeedStore

func SetFeedStore(store *store.FeedStore) {
	feedStore = store
}

func Subscribe(feedUrl string) error {
	s := *feedStore
	return s.Save(feedUrl)
}

func Unsubscribe(feedUrl string) error {
	s := *feedStore
	return s.Delete(feedUrl)
}
