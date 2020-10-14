package router

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/YuhriBernardes/rest_go/router/route"
	"github.com/gorilla/mux"
)

type MuxRouter struct {
	routes      []route.RouteConfig
	middlewares []mux.MiddlewareFunc
}

func (router *MuxRouter) Name() string {
	return "gorilla/mux"
}

func (router *MuxRouter) RegisterRoute(routeConfig route.RouteConfig) error {

	if router.routes == nil {
		router.routes = make([]route.RouteConfig, 0)
	}

	log.WithFields(log.Fields{
		"path":   routeConfig.Path,
		"method": routeConfig.Method,
	}).Info("Route registry")

	router.routes = append(router.routes, routeConfig)
	return nil
}

func (router *MuxRouter) Create() http.Handler {
	muxRouter := mux.NewRouter()

	for _, route := range router.routes {
		log.WithFields(log.Fields{
			"path":       route.Path,
			"method":     route.Method,
			"routerType": router.Name(),
		}).Warn("Adding route to router")

		muxRouter.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return muxRouter
}
