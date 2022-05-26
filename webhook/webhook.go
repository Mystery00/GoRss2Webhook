package webhook

import (
	"GoRss2Webhook/webhook/action"
	"GoRss2Webhook/webhook/store"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
)

func DoWebhook(webhook store.Webhook, item gofeed.Item) {
	webhookAction := action.GetWebhookAction(webhook.Type)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logrus.Warnf(`do webhook failed, err: %s`, err)
				return
			}
		}()
		webhookAction.DoWebhook(webhook.MetaData, item)
	}()
}
