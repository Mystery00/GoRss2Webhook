package action

import (
	"GoRss2Webhook/webhook/action/http"
	"github.com/mmcdole/gofeed"
)

const (
	HTTP = iota + 1
	TELEGRAM_BOT
	WECOM_BOT
)

type Action interface {
	// DoWebhook 执行webhook的逻辑
	DoWebhook(metaData string, item gofeed.Item)
}

func GetWebhookAction(webhookType int8) Action {
	switch webhookType {
	case HTTP:
		return http.HttpAction{}
	case TELEGRAM_BOT:
	case WECOM_BOT:

	}
	return nil
}
