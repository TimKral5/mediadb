package mediadb

import (
	"fmt"
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

var levelCode = map[LogLevel]string{
	debugLevel:   "\033[34mDBG",
	logLevel:     "\033[37mLOG",
	infoLevel:    "\033[036mINF",
	warningLevel: "\033[33mWRN",
	errorLevel:   "\033[31mERR",
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
	return levelCode[lvl]
}

func (logger *Logger) LogByLevel(level LogLevel, args ...any) {
	ctime := time.Now().Format(time.DateTime)

	fmt.Print(ctime, " | ", level.String(), ":\033[0m ")
	fmt.Println(args...)
}

// If debug is enabled, a message will be logged.
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

func (logger *Logger) Warn(args ...any) {
	logger.LogByLevel(warningLevel, args...)
}

func (logger *Logger) Error(args ...any) {
	logger.LogByLevel(errorLevel, args...)
}
