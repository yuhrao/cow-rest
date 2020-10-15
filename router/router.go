package router

import (
	"net/http"

	"github.com/YuhriBernardes/rest_go/middleware"
	"github.com/YuhriBernardes/rest_go/route"
)

type Router interface {
	RegisterRoute(r route.RouteConfig) error
	RegisterMiddleware(handler middleware.MuxMiddleware)
	Create() http.Handler
	Name() string
}
