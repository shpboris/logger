package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

const (
	logLevelKey     = "LOG_LEVEL"
	logFilePathKey  = "LOG_FILE_PATH"
	reportCallerKey = "REPORT_CALLER"
	filePermissions = 0666
)

var Log = logrus.New()

func init() {
	Log.SetLevel(getLogLevel())
	f, err := os.OpenFile(getLogFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, filePermissions)
	if err != nil {
		Log.Fatalf("Error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)
	Log.SetOutput(wrt)
	Log.SetReportCaller(getReportCaller())
}

func getLogLevel() logrus.Level {
	lvl, ok := os.LookupEnv(logLevelKey)
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	return ll
}

func getLogFilePath() string {
	path, ok := os.LookupEnv(logFilePathKey)
	if !ok {
		path = "app.log"
	}
	return path
}

func getReportCaller() bool {
	reportCaller, ok := os.LookupEnv(reportCallerKey)
	if !ok || strings.EqualFold(reportCaller, "y") {
		return true
	}
	return false
}
