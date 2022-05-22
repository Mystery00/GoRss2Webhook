package viper

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(parent, name, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType(configType)
	v.AddConfigPath(parent)
	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatal("read config failed", err)
	}
	return v
}

func WatchChange(v *viper.Viper, run func(v *viper.Viper)) {
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		logrus.Info("config file updated")
		err := v.ReadInConfig()
		if err != nil {
			logrus.Fatal("read config failed", err)
		}
		run(v)
	})
}
