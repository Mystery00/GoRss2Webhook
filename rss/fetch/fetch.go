package fetch

import (
	"GoRss2Webhook/rss/store"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

type ParserClient struct {
	*gofeed.Parser
	UserAgent  string
	HttpClient *http.Client
}

func Parse(subscriber store.FeedSubscriber) (*gofeed.Feed, error) {
	parserClient := getClient(subscriber)
	logrus.Debugf(`start load rss [%s]`, subscriber.FeedUrl)
	feed, err := parserClient.ParseURL(subscriber.FeedUrl)
	return feed, err
}

var parserClientMap = make(map[string]*ParserClient)

func getClient(subscriber store.FeedSubscriber) *ParserClient {
	client, exist := parserClientMap[subscriber.FeedUrl]
	if exist {
		return client
	}
	ua := subscriber.UserAgent
	if ua == "" {
		ua = "GoRss2Webhook"
	}
	proxy := http.ProxyFromEnvironment
	if subscriber.ProxyUrl != "" {
		p, _ := url.Parse(subscriber.ProxyUrl)
		proxy = http.ProxyURL(p)
	}
	timeout := subscriber.Timeout
	if timeout <= 0 {
		timeout = time.Minute
	}
	parser := gofeed.NewParser()
	parser.UserAgent = ua
	client = &ParserClient{
		Parser:    parser,
		UserAgent: ua,
		HttpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: proxy,
			},
			Timeout: timeout,
		},
	}
	return client
}
