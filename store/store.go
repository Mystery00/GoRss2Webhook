package store

import (
	"github.com/mmcdole/gofeed"
)

type RssStore interface {
	// Save 保存RSS记录
	Save(feed gofeed.Feed, item gofeed.Item) error

	// Exist 判断有没有历史的RSS记录
	Exist(feed gofeed.Feed, item gofeed.Item) bool

	// Delete 删除RSS记录
	Delete(feed gofeed.Feed) error

	// Clear 清空历史记录
	Clear() error
}

var rssStore *RssStore

func SetRssStore(store *RssStore) {
	rssStore = store
}
