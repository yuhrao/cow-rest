package main

import (
	"sync"

	env "github.com/YuhriBernardes/rest_go/environment"
	"github.com/YuhriBernardes/rest_go/middleware"
	"github.com/YuhriBernardes/rest_go/router"
	"github.com/YuhriBernardes/rest_go/router/route"
	"github.com/YuhriBernardes/rest_go/server"
	log "github.com/sirupsen/logrus"
)

var (
	once = sync.Once{}
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

	routerImpl.RegisterMiddleware(middleware.RequestLogger{})

	srvr := &server.Server{}

	srvr.Init(3000, routerImpl)

	if err := srvr.Start(); err != nil {
		log.WithField("error", err).Error("Failed to start server")
	}
}
