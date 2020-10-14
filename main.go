package main

import (
	"net/http"
	"strconv"

	env "github.com/YuhriBernardes/rest_go/environment"
	"github.com/YuhriBernardes/rest_go/router"
	"github.com/YuhriBernardes/rest_go/router/route"
	log "github.com/sirupsen/logrus"
)

func startLogger() {

	if appEnv := env.GetString("APP_ENV", env.DEVELOPMENT); appEnv == env.DEVELOPMENT {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
		})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	startLogger()

	routerImpl := &router.MuxRouter{}
	routerImpl.RegisterRoute(route.HealthCheckRoute)

	server := http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(3000),
		Handler: routerImpl.Create(),
	}

	log.WithFields(log.Fields{
		"address": server.Addr,
		"router":  routerImpl.Name(),
	}).Warn("Starting server")

	if err := server.ListenAndServe(); err != nil {
		log.WithField("error", err).Error("Failed to start server")
	}

}
