package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/thanhpv3380/ftl-common-go/utils"
)

var Logger *logrus.Logger

type LoggerConfig struct {
	LogFile          string       // Tên tệp log
	LogLevel         logrus.Level // Mức độ log
	TimestampFormat  string       // Định dạng timestamp
	DisableColors    bool         // Bật tắt màu sắc
	FullTimestamp    bool         // Hiển thị timestamp đầy đủ
	ForceColors      bool         // Bắt buộc sử dụng màu sắc (nếu cần)
	PadLevelText     bool         // Thêm khoảng trắng cho text mức độ log
	QuoteEmptyFields bool         // Trích dẫn các trường rỗng
	DisableQuote     bool         // Không trích dẫn chuỗi
}

func DefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		LogFile:          "logs/app.log",
		LogLevel:         logrus.InfoLevel,
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.000",
		ForceColors:      true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
		DisableQuote:     true,
	}
}

func InitLogger(config *LoggerConfig) {
	if config == nil {
		config = DefaultLoggerConfig()
	}

	Logger = logrus.New()

	logDir := filepath.Dir(config.LogFile)
	utils.CheckExistOrMake(logDir)

	file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatal("Unable to open log file:", err)
	}

	Logger.SetOutput(io.MultiWriter(os.Stdout, file))

	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:    config.DisableColors,
		FullTimestamp:    config.FullTimestamp,
		TimestampFormat:  config.TimestampFormat,
		ForceColors:      config.ForceColors,
		PadLevelText:     config.PadLevelText,
		QuoteEmptyFields: config.QuoteEmptyFields,
		DisableQuote:     config.DisableQuote,
	})
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

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
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
