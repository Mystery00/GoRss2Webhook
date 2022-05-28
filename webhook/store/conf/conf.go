package conf

import (
	"GoRss2Webhook/utils/conf"
	"GoRss2Webhook/webhook/action"
	"GoRss2Webhook/webhook/store"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type confStore struct {
	storePath  string
	fileName   string
	configType string
	data       map[string][]store.Webhook
}

type webhookData struct {
	Type     string                 `mapstructure:"type"`
	MetaData map[string]interface{} `mapstructure:"metaData"`
}

func Init(storePath, fileName, configType string) store.WebhookStore {
	var m map[string][]webhookData
	c, err := conf.UnmarshalFromFile(&m, storePath, fileName, configType)
	if err != nil {
		logrus.Warnf(`parse conf error: %s`, err.Error())
	}
	var feedStore store.WebhookStore
	configStore := &confStore{
		storePath:  storePath,
		fileName:   fileName,
		configType: configType,
		data:       parseData(m),
	}
	conf.WatchChange(c, func(v *viper.Viper) {
		var m map[string][]webhookData
		_, err := conf.UnmarshalFromFile(&m, storePath, fileName, configType)
		if err != nil {
			logrus.Warnf(`parse conf error: %s`, err.Error())
		}
		configStore.data = parseData(m)
	})
	feedStore = configStore
	return feedStore
}

func hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func parseData(m map[string][]webhookData) map[string][]store.Webhook {
	mm := make(map[string][]store.Webhook, 0)
	for k, element := range m {
		l := make([]store.Webhook, 0)
		for _, data := range element {
			bytes, err := json.Marshal(data.MetaData)
			if err != nil {
				logrus.Warnf(`parse conf error: %s`, err.Error())
				continue
			}
			webhookType := action.HTTP
			switch data.Type {
			case "http":
				webhookType = action.HTTP
				break
			case "telegram":
				webhookType = action.TELEGRAM_BOT
				break
			case "wecom":
				webhookType = action.WECOM_BOT
				break
			}
			ww := store.Webhook{
				Type:     webhookType,
				MetaData: string(bytes),
			}
			l = append(l, ww)
		}
		mm[k] = l
	}
	return mm
}

func (store *confStore) Save(feedUrl string, webhook store.Webhook) error {
	u := hash(feedUrl)
	store.data[u] = append(store.data[u], webhook)
	return conf.MarshalToFile(store.data, store.storePath, store.fileName, store.configType)
}

func (store *confStore) GetAll(feedUrl string) ([]store.Webhook, error) {
	u := hash(feedUrl)
	if feedUrl == "common" {
		u = feedUrl
	}
	return store.data[u], nil
}

func (store *confStore) Delete(feedUrl string) error {
	u := hash(feedUrl)
	delete(store.data, u)
	return conf.MarshalToFile(store.data, store.storePath, store.fileName, store.configType)
}
