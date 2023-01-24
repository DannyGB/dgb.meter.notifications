package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dgb/meter.notifications/internal/configuration"

	"github.com/gorilla/mux"
)

func HandleRequests(conf configuration.Configuration) {

	myRouter := mux.NewRouter().StrictSlash(true)
	subRoute := myRouter.PathPrefix("/api").Subrouter()

	subRoute.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("OK")
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.HTTP_PORT), subRoute))
}
