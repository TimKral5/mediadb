package mediadb

import (
	"fmt"
	"strings"
	"time"
)

type LogLevel int

const (
	debugLevel LogLevel = iota
	logLevel
	infoLevel
	warningLevel
	errorLevel
)

var levelName = map[LogLevel]string{
	debugLevel:   "debug",
	logLevel:    "log",
	infoLevel:    "info",
	warningLevel: "warning",
	errorLevel:   "error",
}

type Logger struct {
	EnableDebug bool
}

func NewLogger() Logger {
	logger := Logger{
		EnableDebug: false,
	}
	return logger
}

func (lvl LogLevel) String() string {
	return levelName[lvl]
}

func (logger *Logger) LogByLevel(level LogLevel, args ...any) {
	ctime := time.Now().Format(time.DateTime)
	prefix := strings.ToUpper(level.String())

	fmt.Print(ctime, " | ", prefix, ": ")
	fmt.Println(args...)
}

func (logger *Logger) Debug(args ...any) {
	if !logger.EnableDebug {
		return
	}

	logger.LogByLevel(debugLevel, args...)
}

func (logger *Logger) Log(args ...any) {
	logger.LogByLevel(logLevel, args...)
}

func (logger *Logger) Info(args ...any) {
	logger.LogByLevel(infoLevel, args...)
}
