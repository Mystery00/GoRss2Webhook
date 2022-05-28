package core

import (
	"GoRss2Webhook/config"
	"GoRss2Webhook/feed/fetch"
	feed "GoRss2Webhook/feed/store"
	rss "GoRss2Webhook/store"
	"GoRss2Webhook/webhook"
	"GoRss2Webhook/webhook/store"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

var feedStore *feed.FeedStore
var rssStore *rss.RssStore
var webhookStore *store.WebhookStore

func FeedStore() *feed.FeedStore {
	return feedStore
}

func WebhookStore() *store.WebhookStore {
	return webhookStore
}

func Init() {
	feedStore = getFeedStore()
	rssStore = getRssStore()
	webhookStore = getWebhookStore()
	if viper.GetBool(config.CronEnable) {
		crontab := viper.GetString(config.CronTab)
		//注册定时任务
		c := cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(CronLogger{logrus.StandardLogger()})))
		_, err := c.AddFunc(crontab, func() {
			doWork(*feedStore, *rssStore, *webhookStore)
		})
		if err != nil {
			log.Fatal(`register cron failed`, err)
		}
	}
}

func DoWork() {
	doWork(*feedStore, *rssStore, *webhookStore)
}

func doWork(feedStore feed.FeedStore, rssStore rss.RssStore, webhookStore store.WebhookStore) {
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
