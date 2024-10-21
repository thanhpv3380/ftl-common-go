package logger

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
}

func CreateFields(fields map[string]interface{}) logrus.Fields {
	return logrus.Fields(fields)
}

func Info(message string, fields ...logrus.Fields) {
	if len(fields) == 0 || fields[0] == nil {
		Logger.Info(message)
	} else {
		Logger.WithFields(fields[0]).Info(message)
	}
}

func Warn(message string, fields ...logrus.Fields) {
	if len(fields) == 0 || fields[0] == nil {
		Logger.Warn(message)
	} else {
		Logger.WithFields(fields[0]).Warn(message)
	}
}

func Error(message string, err error, fields ...logrus.Fields) {
	if err == nil {
		Logger.Error(message)
		return
	}

	if len(fields) == 0 || fields[0] == nil {
		Logger.Error(message + ": " + err.Error())
	} else {
		Logger.WithFields(fields[0]).Error(message + ": " + err.Error())
	}
}
