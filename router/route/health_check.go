package route

import (
	"encoding/json"
	"net/http"
)

var HealthCheckRoute = RouteConfig{
	Path:   "/api/check",
	Method: http.MethodGet,
	Handler: func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(200)

		json.NewEncoder(w).Encode(map[string]bool{"ok": true})

	},
}
