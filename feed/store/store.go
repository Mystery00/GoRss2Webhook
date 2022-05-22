package store

import "time"

type FeedSubscriber struct {
	FeedUrl   string
	UserAgent string
	ProxyUrl  string
	Timeout   time.Duration
}

type FeedStore interface {
	// Subscribe 订阅信息
	Subscribe(subscriber FeedSubscriber) error

	// GetAll 获取订阅信息
	GetAll() ([]FeedSubscriber, error)

	// Unsubscribe 取消订阅信息
	Unsubscribe(feedUrl string) error
}
