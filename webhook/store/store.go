package store

import (
	"time"
)

type Webhook struct {
	SubscribeUrl string
	Http         Http
	ProxyUrl     string
	Timeout      time.Duration
}

type Http struct {
	Url    string
	Method string
	Body   string
	Header map[string]string
}

type WebhookStore interface {
	// Save 保存Webhook信息
	Save(webhook Webhook) error

	// GetAll 获取Webhook信息
	GetAll(feedUrl string) ([]Webhook, error)

	// Delete 删除Webhook信息
	Delete(feedUrl string) error
}
