package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger logrus.FieldLogger
}

var logrusLogger = logrus.New()
var Log = NewLogger(logrusLogger)

func NewLogger(logger logrus.FieldLogger) *Logger {
	return &Logger{logger}
}

func (l *Logger) Info(data string) {
	l.logger.Infof(data + "\n")
}

func (l *Logger) Fatal(data string) {
	l.logger.Fatalf(data + "\n")
}

func (l *Logger) Error(data string) {
	l.logger.Error(data + "\n")
}
