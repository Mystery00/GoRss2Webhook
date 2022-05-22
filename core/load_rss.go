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
			feedStore := memory.Init()
			return &feedStore
		}
	case "file":
		{
			storePath := viper.GetString(StoreRssFilePath)
			feedStore := file.Init(storePath)
			return &feedStore
		}
	case "viper":
		{

		}
	}
	feedStore := memory.Init()
	return &feedStore
}
