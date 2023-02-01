package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	log *logrus.Entry
}

func GetLogger(component string) Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return Logger{
		logrus.WithField("component", component),
	}
}

func (l Logger) Info(o interface{}) {
	l.log.Info(o)
}

func (l Logger) Error(o interface{}) {
	l.log.Error(o)
}

func (l Logger) Warn(o interface{}) {
	l.log.Warn(o)
}

func (l Logger) Fatal(o interface{}) {
	l.log.Fatal(o)
}

func (l Logger) Debug(o interface{}) {
	l.log.Debug(o)
}
