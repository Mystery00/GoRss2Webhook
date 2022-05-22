package core

import "github.com/sirupsen/logrus"

type CronLogger struct {
	*logrus.Logger
}

func (cLogger CronLogger) Info(msg string, keysAndValues ...interface{}) {
	cLogger.Infof(msg, keysAndValues...)
}

func (cLogger CronLogger) Error(err error, _ string, _ ...interface{}) {
	cLogger.Errorln(err)
}
