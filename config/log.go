package config

import (
	"fmt"
	"github.com/Mystery00/lumberjack"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

func getLogFilePath(fileName string) string {
	logFileHome, exist := os.LookupEnv(EnvLogHome)
	if !exist {
		logFileHome = viper.GetString(LogHome)
	}
	err := os.MkdirAll(logFileHome, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}
	return fmt.Sprintf(`%s/%s`, logFileHome, fileName)
}

func InitLog() {
	logFile := viper.GetString(LogFile)
	fileName := getLogFilePath(logFile)
	local := viper.GetBool(LogLocal)
	var out io.Writer
	if local {
		//本地启动，日志打印到控制台
		out = os.Stdout
	} else {
		//服务器启动，日志打印到文件中
		out = &lumberjack.Logger{
			Filename:         fileName,
			MaxSize:          256,
			MaxAge:           3,
			LocalTime:        true,
			Compress:         true,
			SplitByDay:       true,
			BackupTimeFormat: `2006-01-02.150405`,
		}
	}
	//设置输出
	logrus.SetOutput(out)
	//设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        true,
		NoColors:        !viper.GetBool(LogColor),
	})
}
