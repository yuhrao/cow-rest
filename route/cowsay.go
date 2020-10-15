package route

import (
	"fmt"
	"net/http"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/gorilla/mux"
)

var (
	GetCows = RouteConfig{
		Path:   "/api/cows",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request, b interface{}) (statusCode int, responsebody interface{}, err error) {
			cowList := make([]string, 0)
			originalCowlist, _ := cowsay.AssetDir("cows")
			for _, cowName := range originalCowlist {
				formattedName := strings.ReplaceAll(cowName, ".cow", "")
				cowList = append(cowList, formattedName)
			}
			return 200, cowList, nil
		},
	}
	GetCowsay = RouteConfig{
		Path:   "/api/cows/{cowname}",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request, b interface{}) (statusCode int, responsebody interface{}, err error) {
			pathVars := mux.Vars(r)
			cowName, found := pathVars["cowname"]

			if !found {
				return 404, map[string]string{"error": "Cow name is required"}, nil
			}

			if !cowExists(cowName) {
				return 404, map[string]string{"error": fmt.Sprintf("Cow %s not found", cowName)}, nil
			}
			phrase := r.URL.Query().Get("text")

			say, _ := cowsay.Say(
				cowsay.Phrase(phrase),
				cowsay.Type(cowName),
			)

			w.Header().Set("Content-Type", "text/plain")

			return 200, say, nil
		},
	}
)

func cowExists(cowName string) bool {
	internalCowNames, _ := cowsay.AssetDir("cows")
	for _, internalCowName := range internalCowNames {
		if internalCowName == cowName+".cow" {
			return true
		}
	}

	return false
}
