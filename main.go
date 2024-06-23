package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", AddExoplanetHandler).Methods("POST")
	r.HandleFunc("/exoplanets", ListExoplanetsHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", GetExoplanetHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", UpdateExoplanetHandler).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", DeleteExoplanetHandler).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel-estimation", FuelEstimationHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8880", r))
}
