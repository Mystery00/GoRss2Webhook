package core

import (
	"GoRss2Webhook/rss/fetch"
	feed "GoRss2Webhook/rss/store"
	rss "GoRss2Webhook/store"
	"GoRss2Webhook/webhook"
	"GoRss2Webhook/webhook/store"
	"github.com/sirupsen/logrus"
)

func DoWork(feedStore feed.FeedStore, rssStore rss.RssStore, webhookStore store.WebhookStore) {
	subscribers, err := feedStore.GetAll()
	if err != nil {
		panic(err)
	}
	for _, subscriber := range subscribers {
		subscriber := subscriber
		go func() {
			defer func() {
				if err := recover(); err != nil {
					logrus.Warnf(`load rss failed, err: %s`, err)
					return
				}
			}()
			resultFeed, err := fetch.Parse(subscriber)
			if err != nil {
				panic(err)
			}
			for _, item := range resultFeed.Items {
				exist := rssStore.Exist(*resultFeed, *item)
				if !exist {
					logrus.Debugf(`save feed [%s] item [%s]`, resultFeed.Link, item.GUID)
					err := rssStore.Save(*resultFeed, *item)
					if err != nil {
						logrus.Panicf(`save feed [%s] item [%s] failed`, resultFeed.Link, item.GUID)
					}
					webhooks, err := webhookStore.GetAll(subscriber.FeedUrl)
					if err != nil {
						logrus.Panicf(`get webhook list [%s] failed`, subscriber.FeedUrl)
					}
					for _, w := range webhooks {
						logrus.Debugf(`do webhook %s`, w)
						webhook.DoWebhook(w, *item)
					}
					if len(webhooks) == 0 {
						logrus.Warnf(`parse rss [%s] success but no webhook config`, subscriber.FeedUrl)
					}
				} else {
					logrus.Debugf(`item [%s] already exist`, item.GUID)
				}
			}
		}()
	}
}
