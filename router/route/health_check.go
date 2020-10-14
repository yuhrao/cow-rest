package route

import (
	"net/http"
)

var HealthCheckRoute = RouteConfig{
	Path:   "/api/check",
	Method: http.MethodGet,
	Handler: func(w http.ResponseWriter, r *http.Request, _ interface{}) (statusCode int, responseBody interface{}, err error) {

		return http.StatusAccepted, map[string]bool{"ok": true}, nil

	},
}
