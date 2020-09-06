package requests

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	baseLogger := logrus.New()
	logger := Logger{baseLogger}
	var err error
	// set REQUESTS_LOGLEVEL for log level, defaults to info
	level, exist := os.LookupEnv("REQUESTS_LOGLEVEL")
	if !exist{
		level = "info"
	}
	logger.Level, err = logrus.ParseLevel(level)
	if err != nil{
		panic(err)
	}
	// setting logger format to string
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// set to true for showing filename and line number from where logger being called
	logger.SetReportCaller(true)
	return &logger
}
