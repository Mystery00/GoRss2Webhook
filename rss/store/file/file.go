package file

import (
	"GoRss2Webhook/rss/store"
	"encoding/json"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type fileStore struct {
	storePath string
	fileName  string
	data      []store.FeedSubscriber
}

func Init(storePath, fileName, saveCron string) store.FeedStore {
	data := make([]store.FeedSubscriber, 0)
	bytes, err := readFile(storePath, fileName)
	if err != nil {
		logrus.Warnf(`read file error: %s`, err.Error())
	} else if bytes != nil {
		err := json.Unmarshal(bytes, &data)
		if err != nil {
			logrus.Warnf(`unmarshal file error: %s`, err.Error())
		} else {
			logrus.Infof(`read exist file success`)
		}
	}
	var rssStore store.FeedStore
	fileStore := &fileStore{
		storePath: storePath,
		fileName:  fileName,
		data:      data,
	}
	//注册定时任务
	cronJob := cron.New(cron.WithSeconds())
	_, err = cronJob.AddFunc(saveCron, func() {
		//每10秒执行一次
		err := fileStore.doSave()
		if err != nil {
			logrus.Warnf(`save file error: %s`, err.Error())
		}
	})
	if err != nil {
		logrus.Warnf(`add cron job error: %s`, err.Error())
	}
	rssStore = fileStore
	return rssStore
}

func (store *fileStore) doSave() error {
	bytes, err := json.Marshal(store.data)
	if err != nil {
		return err
	}
	return writeFile(store.storePath, store.fileName, bytes)
}

func readFile(storePath, fileName string) ([]byte, error) {
	filePath := storePath + "/" + fileName
	if !exists(filePath) {
		return nil, nil
	}
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// 写入文件
func writeFile(parent, fileName string, content []byte) error {
	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := parent + "/" + fileName
	err = ioutil.WriteFile(filePath, content, 0644)
	return err
}

func (store *fileStore) Save(subscriber store.FeedSubscriber) error {
	store.data = append(store.data, subscriber)
	return nil
}

func (store *fileStore) GetAll() ([]store.FeedSubscriber, error) {
	return store.data, nil
}

func (store *fileStore) Delete(feedUrl string) error {
	for i, subscriber := range store.data {
		if subscriber.FeedUrl == feedUrl {
			store.data = append(store.data[:i], store.data[i+1:]...)
			return nil
		}
	}
	return nil
}

func exists(path string) bool {
	//获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
