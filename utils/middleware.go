package utils

import "net/http"

type Middleware func(next http.Handler) http.Handler

// Creates a stack of middleware that is chained sequencially from
// the last through to the first handler.
func CreateStack(arr ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(arr) - 1; i >= 0; i-- {
			middleware := arr[i]
			next = middleware(next)
		}

		return next
	}
}
