package middleware

import "github.com/gorilla/mux"

type MuxMiddleware interface {
	Name() string
	Get() mux.MiddlewareFunc
}
