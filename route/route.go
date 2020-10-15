package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request, entity interface{}) (statusCode int, responseBody interface{}, err error)

type RouteConfig struct {
	Path      string
	Handler   RouteHandler
	Method    string
	GetEntity func() interface{}
}

func readJson(r *http.Request, entity *interface{}) {
	if r.Body == nil {
		return
	}

	jsonBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("Failed to read JSON")

	}
	json.Unmarshal(jsonBytes, entity)
	return
}

func Handler(route RouteConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entity interface{}

		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			entity := route.GetEntity()
			readJson(r, &entity)
		}

		statusCode, response, _ := route.Handler(w, r, entity)

		contentType := w.Header().Get("Content-Type")
		if contentType == "" {
			contentType = "application/json"
			w.Header().Add("Content-Type", "application/json")
		}

		if contentType == "application/json" {

			w.WriteHeader(statusCode)

			json.NewEncoder(w).Encode(response)

		} else if contentType == "text/plain" {
			w.WriteHeader(statusCode)
			w.Write([]byte(response.(string)))
		} else {
			w.WriteHeader(statusCode)
		}

	}
}
