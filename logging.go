package main

import (
	mediadb "mediadb/internals"
	"net/http"
	"strconv"
	"time"
)

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
func EnableLogging(log mediadb.Logger, next http.Handler) http.Handler {
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
