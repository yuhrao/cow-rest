package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type RequestLogger struct {
}

func (md RequestLogger) Name() string {
	return "request_logger"
}

func (md RequestLogger) Get() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			log.WithFields(log.Fields{
				"reqUri": r.RequestURI,
				"method": r.Method,
			}).Info("New request")

			h.ServeHTTP(w, r)
		})
	}
}
