package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	exoplanets = make(map[string]*Exoplanet)
	idCounter  = 1
	mutex      = &sync.Mutex{}
)

func AddExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	var exoplanet Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := exoplanet.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	exoplanet.ID = strconv.Itoa(idCounter)
	idCounter++
	exoplanets[exoplanet.ID] = &exoplanet
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func ListExoplanetsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	var exoplanetList []*Exoplanet
	for _, exo := range exoplanets {
		exoplanetList = append(exoplanetList, exo)
	}

	mutex.Unlock()
	json.NewEncoder(w).Encode(exoplanetList)
}

func GetExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	mutex.Lock()
	exoplanet, exists := exoplanets[id]
	mutex.Unlock()
	if !exists {
		http.Error(w, "exoplanet not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	mutex.Lock()
	_, exist := exoplanets[id]
	mutex.Unlock()
	if !exist {
		http.Error(w, "not available!", http.StatusBadRequest)
		return
	}

	var UpdatedExoplanet Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&UpdatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := UpdatedExoplanet.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UpdatedExoplanet.ID = id
	mutex.Lock()
	exoplanets[id] = &UpdatedExoplanet
	mutex.Unlock()

	json.NewEncoder(w).Encode(UpdatedExoplanet)
}

func DeleteExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mutex.Lock()
	_, exists := exoplanets[id]
	if !exists {
		http.Error(w, "not exists!", http.StatusBadRequest)
	}

	if exists {
		delete(exoplanets, id)
	}
	mutex.Unlock()

	w.WriteHeader(http.StatusAccepted)
}

func FuelEstimationHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req FuelEstimationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	exoplante, exist := exoplanets[id]
	if !exist {
		http.Error(w, "not exists!", http.StatusBadRequest)
		return
	}

	var gravity float64
	switch exoplante.Type {
	case GasGaint:
		gravity = 0.5 / (exoplante.Radius * exoplante.Radius)
	case Terrestrial:
		gravity = exoplante.Mass / (exoplante.Radius * exoplante.Radius)
	default:
		http.Error(w, "unknown type", http.StatusBadRequest)
		return
	}

	fuel := float64(exoplante.Distance) / gravity * gravity * float64(req.CrewCapacity)
	resp := FuelEstimationResponse{FuelUnits: fuel}
	json.NewEncoder(w).Encode(resp)
}
