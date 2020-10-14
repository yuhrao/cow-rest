package route

import "net/http"

type RouteConfig struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}
