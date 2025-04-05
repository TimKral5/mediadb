package main

import (
	mediadb "mediadb/internals"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func EnableLogging(log mediadb.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode: http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		log.Log(wrapped.statusCode, r.Method, r.URL, time.Since(start))
	})
}
