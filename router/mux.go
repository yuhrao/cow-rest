package router

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/YuhriBernardes/rest_go/middleware"
	"github.com/YuhriBernardes/rest_go/router/route"
	"github.com/gorilla/mux"
)

type MuxRouter struct {
	routes      []route.RouteConfig
	middlewares []middleware.MuxMiddleware
}

func (router *MuxRouter) RegisterMiddleware(muxMiddleware middleware.MuxMiddleware) {
	if router.middlewares == nil {
		router.middlewares = make([]middleware.MuxMiddleware, 0)
	}

	log.WithField("middleware", muxMiddleware.Name()).Info("Adding middleware")

	router.middlewares = append(router.middlewares, muxMiddleware)
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

	if router.middlewares == nil {
		log.Warning("No middlewares to register")
	} else {
		for _, middleware := range router.middlewares {

			log.WithField("name", middleware.Name()).Warn("Adding middleware to router")

			muxRouter.Use(middleware.Get())
		}
	}

	for _, routeConfig := range router.routes {
		log.WithFields(log.Fields{
			"path":       routeConfig.Path,
			"method":     routeConfig.Method,
			"routerType": router.Name(),
		}).Warn("Adding route to router")

		muxRouter.HandleFunc(routeConfig.Path, route.Handler(routeConfig)).Methods(routeConfig.Method)
	}

	return muxRouter
}
