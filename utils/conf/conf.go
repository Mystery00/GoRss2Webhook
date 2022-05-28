package conf

import (
	"GoRss2Webhook/utils/file"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func InitConfig(parent, name, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType(configType)
	v.AddConfigPath(parent)
	doRead(v)
	return v
}

func WatchChange(v *viper.Viper, run func(v *viper.Viper)) {
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		logrus.Info("conf file updated")
		doRead(v)
		run(v)
	})
}

func doRead(v *viper.Viper) {
	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatal("read conf failed", err)
	}
}

func UnmarshalFromFile(rawVal interface{}, parent, fileName, configType string) (v *viper.Viper, err error) {
	if !file.Exists(parent) {
		err := os.MkdirAll(parent, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	filePath := fmt.Sprintf(`%s/%s.%s`, parent, fileName, configType)
	if !file.Exists(filePath) {
		_, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		v := viper.New()
		v.SetConfigName(fileName)
		v.SetConfigType(configType)
		v.AddConfigPath(parent)
		return v, nil
	}
	v = InitConfig(parent, fileName, configType)
	return v, v.Unmarshal(rawVal)
}

func MarshalToFile(values interface{}, storePath, fileName, configType string) error {
	var bytes []byte
	switch configType {
	case "yaml":
		bs, err := yaml.Marshal(values)
		if err != nil {
			return err
		}
		bytes = bs
		break
	case "json":
		bs, err := json.Marshal(values)
		if err != nil {
			return err
		}
		bytes = bs
		break
	}
	name := fmt.Sprintf(`%s.%s`, fileName, configType)
	return writeFile(storePath, name, bytes)
}

func writeFile(parent, fileName string, content []byte) error {
	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := parent + "/" + fileName
	err = ioutil.WriteFile(filePath, content, 0644)
	return err
}
