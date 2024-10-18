package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// InitLogger initializes the logger with a specific level
func InitLogger(level logrus.Level) {
	log.SetLevel(level)
	log.SetFormatter(&logrus.JSONFormatter{})
}

// Info logs an informational message
func Info(message string) {
	log.Info(message)
}

// Error logs an error message
func Error(message string) {
	log.Error(message)
}

// WithFields logs a message with custom fields
func WithFields(fields logrus.Fields, message string) {
	log.WithFields(fields).Info(message)
}
