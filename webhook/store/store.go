package store

type Webhook struct {
	SubscribeUrl string
	Type         int8
	MetaData     string
}

type WebhookStore interface {
	// Save 保存Webhook信息
	Save(webhook Webhook) error

	// GetAll 获取Webhook信息
	GetAll(feedUrl string) ([]Webhook, error)

	// Delete 删除Webhook信息
	Delete(feedUrl string) error
}
