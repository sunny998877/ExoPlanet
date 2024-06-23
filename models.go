package main

import "errors"

type ExoplanetType string

const (
	GasGaint    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"mame"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        float64       `json:"mass"`
	Type        ExoplanetType `json:"type"`
}

type FuelEstimationRequest struct {
	CrewCapacity int `json:"crew_capacity"`
}

type FuelEstimationResponse struct {
	FuelUnits float64 `json:"fule_units"`
}

func (e *Exoplanet) Validate() error {
	if e.Name == "" || e.Description == "" || e.Distance < 5 || e.Distance > 1000 || e.Radius < 0.1 || e.Radius > 10 {
		return errors.New("invalid exoplanet data")
	}

	if e.Type == Terrestrial && (e.Mass < 0.1 || e.Mass > 10) {
		return errors.New("invalid mass for terrestrial exoplanet")
	}
	return nil

}
