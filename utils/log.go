package utils

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
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

// A wrapper structure for the writer that stores the status code.
type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func getHttpStatusColorPrefix(statusCode int) string {
	if statusCode >= 100 && statusCode < 200 {
		return "\033[36m"
	} else if statusCode >= 200 && statusCode < 300 {
		return "\033[32m"
	} else if statusCode >= 300 && statusCode < 400 {
		return "\033[35m"
	} else if statusCode >= 400 && statusCode < 600 {
		return "\033[31m"
	}

	return "\033[33m"
}

// A wrapper for an http handler that logs all requests made to the
// input handler to a given logger.
func (log *Logger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		statusPrefix := getHttpStatusColorPrefix(wrapped.statusCode)
		status := statusPrefix + strconv.Itoa(wrapped.statusCode) + "\033[0m"

		log.Log(status, r.Method, r.URL, "\033[33m"+time.Since(start).String()+"\033[0m")
	})
}
