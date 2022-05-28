package core

import (
	"GoRss2Webhook/webhook/store"
	"GoRss2Webhook/webhook/store/conf"
	"GoRss2Webhook/webhook/store/file"
	"GoRss2Webhook/webhook/store/memory"
	"github.com/spf13/viper"
)

func getWebhookStore() *store.WebhookStore {
	feedType := viper.GetString(StoreWebhookType)
	switch feedType {
	case "memory":
		{
			webhookStore := memory.Init()
			return &webhookStore
		}
	case "file":
		{
			storePath := viper.GetString(StoreWebhookFilePath)
			fileName := viper.GetString(StoreWebhookFileName)
			webhookStore := file.Init(storePath, fileName)
			return &webhookStore
		}
	case "viper":
		{
			storePath := viper.GetString(StoreWebhookViperPath)
			fileName := viper.GetString(StoreWebhookViperName)
			viperType := viper.GetString(StoreWebhookViperType)
			if viperType == "" {
				viperType = "yaml"
			}
			webhookStore := conf.Init(storePath, fileName, viperType)
			return &webhookStore
		}
	}
	feedStore := memory.Init()
	return &feedStore
}
