package conf

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type logInfo struct {
	Log         *logrus.Logger
	LogLevel    string `toml:"level" env:"LOG_LEVEL"`
	LogSavePath string `toml:"logfilepath" env:"LOG_FILE_PATH"`
	OsOutput    string `toml:"output" env:"OUTPUT"`
}

func (l *logInfo) InitLogger() error {
	switch l.LogLevel {
	case "info":
		l.Log.SetLevel(logrus.InfoLevel)
	case "debug":
		l.Log.SetLevel(logrus.DebugLevel)
	case "warn":
		l.Log.SetLevel(logrus.WarnLevel)
	default:
		return fmt.Errorf("not supported log level configuration, you can only set the log level as \"debug\" \"info\" or \"warn\", please recheck your environment.toml")
	}
	switch l.OsOutput {
	case "stdOutput":
		l.Log.SetOutput(os.Stdout)
	default:
		return fmt.Errorf("not supported os output configuration, you can only set \"stdOutput\"")
	}
	l.Log.SetReportCaller(true)
	l.Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return nil
}
