package helper

import (
	"fmt"
	"io"
	"os"

	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
)

// SaveLogError save log error to file (THIS LOGGER MOVE TO LOGGER SERVICE)
func SaveLogError(message string, data interface{}) {
	f, _ := os.OpenFile(fmt.Sprintf("logs/%s-error.log", carbon.Now().Format("Y-m-d")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(io.MultiWriter(f, os.Stdout))
	logger.WithField("error", data).Log(logrus.ErrorLevel, message)
}

func SaveErrorProduction(message string, data interface{}) {
	f, _ := os.OpenFile(fmt.Sprintf("logs/%s-error-%s.log", carbon.Now().Format("Y-m-d"), GetConfig("APP_ENV")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(io.MultiWriter(f, os.Stdout))
	logger.WithField("error", data).Log(logrus.ErrorLevel, message)
}
