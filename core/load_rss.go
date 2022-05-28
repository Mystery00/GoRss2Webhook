package core

import (
	"GoRss2Webhook/store"
	"GoRss2Webhook/store/file"
	"GoRss2Webhook/store/memory"
	"github.com/spf13/viper"
)

func getRssStore() *store.RssStore {
	feedType := viper.GetString(StoreRssType)
	switch feedType {
	case "memory":
		{
			rssStore := memory.Init()
			return &rssStore
		}
	case "file":
		{
			storePath := viper.GetString(StoreRssFilePath)
			rssStore := file.Init(storePath)
			return &rssStore
		}
	}
	feedStore := memory.Init()
	return &feedStore
}
