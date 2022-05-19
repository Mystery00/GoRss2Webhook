package store

import "time"

type FeedSubscriber struct {
	FeedUrl   string
	UserAgent string
	ProxyUrl  string
	Timeout   time.Duration
}

type FeedStore interface {
	// Save 保存订阅信息
	Save(subscriber FeedSubscriber) error

	// GetAll 获取订阅信息
	GetAll() ([]FeedSubscriber, error)

	// Delete 删除订阅信息
	Delete(feedUrl string) error
}
