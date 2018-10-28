package logger

import (
	"os"

	"gbs/gream/env"

	"github.com/sirupsen/logrus"
)

type Fields = logrus.Fields

var defaultFields = &Fields{}
var glog = logrus.WithFields(*defaultFields)

func init() {

	logrus.SetOutput(os.Stdout)
	if env.IsDevelopment() {
		// TODO: 这个地方可以重新定义log的格式,开发模式下当打印在终端的时候去掉输出前面的INFO等等
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: true,
		})
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.WarnLevel)
	}

}

func Debug(args ...interface{}) {
	glog.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	glog.Debugf(format, args...)
}

func Info(args ...interface{}) {
	glog.Info(args...)
}

func Warn(args ...interface{}) {
	glog.Warn(args...)
}

func Error(args ...interface{}) {
	glog.Error(args...)
}

func Fatal(args ...interface{}) {
	glog.Fatal(args...)
}

func Panic(args ...interface{}) {
	glog.Panic(args...)
}
