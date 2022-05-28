package core

import (
	"GoRss2Webhook/feed/store"
	"GoRss2Webhook/feed/store/conf"
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
			storePath := viper.GetString(StoreFeedViperPath)
			fileName := viper.GetString(StoreFeedViperName)
			viperType := viper.GetString(StoreFeedViperType)
			if viperType == "" {
				viperType = "yaml"
			}
			feedStore := conf.Init(storePath, fileName, viperType)
			return &feedStore
		}
	}
	feedStore := memory.Init()
	return &feedStore
}
