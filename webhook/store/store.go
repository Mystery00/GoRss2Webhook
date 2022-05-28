package store

type Webhook struct {
	Type     int8   `mapstructure:"type"`
	MetaData string `mapstructure:"metaData"`
}

type WebhookStore interface {
	// Save 保存Webhook信息
	Save(feedUrl string, webhook Webhook) error

	// GetAll 获取Webhook信息
	GetAll(feedUrl string) ([]Webhook, error)

	// Delete 删除Webhook信息
	Delete(feedUrl string) error
}
