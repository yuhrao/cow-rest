package main

import (
	env "github.com/YuhriBernardes/rest_go/environment"
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

	log.Debug("Hello world")
}
