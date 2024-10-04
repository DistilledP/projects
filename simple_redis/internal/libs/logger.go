package libs

import (
	"log"
	"os"
)

const loggerLevelInfo = "INFO"
const loggerLevelWarn = "WARN"
const loggerLevelError = "ERROR"
const loggerLevelDebug = "DEBUG"
const loggerLevelFatal = "FATAL"

var logLevelMap map[string]int

func init() {
	logLevelMap = map[string]int{
		loggerLevelError: 1,
		loggerLevelWarn:  2,
		loggerLevelInfo:  3,
		loggerLevelDebug: 4,
	}
}

type LineLogger struct {
	logger   *log.Logger
	logLevel int
}

func newLogger(prefix string, logLevel int) *LineLogger {
	return &LineLogger{
		logger:   log.New(os.Stdout, prefix, 3),
		logLevel: logLevel,
	}
}

func (l *LineLogger) log(level string, data interface{}) {
	switch level {
	case loggerLevelInfo, loggerLevelDebug, loggerLevelError, loggerLevelWarn:
		l.logger.Printf("[%s]: %v", level, data)
	case loggerLevelFatal:
		l.logger.Fatalf("[%s]: %v", level, data)
	default:
		l.logger.Printf("[%s]: %v", "UNKNOWN", data)
	}
}

func (l *LineLogger) Fatal(data interface{}) {
	l.log(loggerLevelFatal, data)
}

func (l *LineLogger) Error(data interface{}) {
	if l.logLevel >= logLevelMap[loggerLevelError] {
		l.log(loggerLevelError, data)
	}
}

func (l *LineLogger) Info(data interface{}) {
	if l.logLevel >= logLevelMap[loggerLevelInfo] {
		l.log(loggerLevelInfo, data)
	}
}

func (l *LineLogger) Warn(data interface{}) {
	if l.logLevel >= logLevelMap[loggerLevelWarn] {
		l.log(loggerLevelWarn, data)
	}
}

func (l *LineLogger) Debug(data interface{}) {
	if l.logLevel >= logLevelMap[loggerLevelDebug] {
		l.log(loggerLevelDebug, data)
	}
}
