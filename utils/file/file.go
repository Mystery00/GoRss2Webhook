package file

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
)

func Read(parent, fileName string, v any) error {
	bytes, err := readFile(parent, fileName)
	if err != nil {
		panic(err)
	}
	return json.Unmarshal(bytes, v)
}

func Write(parent, fileName string, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return writeFile(parent, fileName, bytes)
}

// Hash 散列名称
func Hash(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 读取文件
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
