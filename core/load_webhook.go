package core

import (
	"GoRss2Webhook/webhook/store"
	"GoRss2Webhook/webhook/store/file"
	"GoRss2Webhook/webhook/store/memory"
	"github.com/spf13/viper"
)

func getWebhookStore() *store.WebhookStore {
	feedType := viper.GetString(StoreWebhookType)
	switch feedType {
	case "memory":
		{
			feedStore := memory.Init()
			return &feedStore
		}
	case "file":
		{
			storePath := viper.GetString(StoreWebhookFilePath)
			fileName := viper.GetString(StoreWebhookFileName)
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
