package file

import (
	"GoRss2Webhook/store"
	"crypto/sha512"
	"encoding/hex"
	"github.com/mmcdole/gofeed"
	"io/ioutil"
	"os"
)

type fileStore struct {
	storePath string
}

func Init(storePath string) store.RssStore {
	var rssStore store.RssStore
	rssStore = &fileStore{
		storePath: storePath,
	}
	return rssStore
}

func hash(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 判断文件是否存在
func existFile(store fileStore, parent, fileName string) bool {
	path := store.storePath + "/" + parent + "/" + fileName
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

// 写入文件
func writeFile(store fileStore, parent, fileName string, content string) error {
	parentPath := store.storePath + "/" + parent
	err := os.MkdirAll(parentPath, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := parentPath + "/" + fileName
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	return err
}

func (store fileStore) Save(feed gofeed.Feed, item gofeed.Item) error {
	parent := hash(feed.Link)
	var fileName string
	if item.GUID != "" {
		fileName = hash(item.GUID)
	} else {
		fileName = hash(item.Link)
	}
	err := writeFile(store, parent, fileName, item.Link)
	return err
}

func (store fileStore) Exist(feed gofeed.Feed, item gofeed.Item) bool {
	parent := hash(feed.Link)
	var fileName string
	if item.GUID != "" {
		fileName = hash(item.GUID)
	} else {
		fileName = hash(item.Link)
	}
	return existFile(store, parent, fileName)
}

func (store fileStore) Delete(feed gofeed.Feed) error {
	key := hash(feed.Link)
	return os.RemoveAll(store.storePath + "/" + key)
}

func (store fileStore) Clear() error {
	return os.RemoveAll(store.storePath)
}
