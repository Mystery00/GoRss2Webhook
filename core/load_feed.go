package core

import (
	"GoRss2Webhook/feed/store"
	"GoRss2Webhook/feed/store/file"
	"GoRss2Webhook/feed/store/memory"
	"github.com/spf13/viper"
)

func getFeedStore() *store.FeedStore {
	feedType := viper.GetString(StoreFeedType)
	switch feedType {
	case "memory":
		{
			feedStore := memory.Init()
			return &feedStore
		}
	case "file":
		{
			storePath := viper.GetString(StoreFeedFilePath)
			fileName := viper.GetString(StoreFeedFileName)
			feedStore := file.Init(storePath, fileName)
			return &feedStore
		}
	case "viper":
		{

		}
	}
	feedStore := memory.Init()
	return &feedStore
}
