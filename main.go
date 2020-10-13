package main

import (
	"net/http"

	env "github.com/YuhriBernardes/rest_go/environment"
	mux "github.com/gorilla/mux"
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

	router := mux.NewRouter()
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World!"))
	}).Methods(http.MethodGet)

	listener := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: router,
	}

	log.WithFields(log.Fields{
		"Host": listener.Addr,
	}).Warn("Starting server")

	listener.ListenAndServe()

}
