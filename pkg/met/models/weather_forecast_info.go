package models

import "time"

type WetherForecastInfo struct {
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type Properties struct {
	Meta       Meta         `json:"meta"`
	Timeseries []Timeseries `json:"timeseries"`
}

type Units struct {
	AirTemperature      string `json:"air_temperature"`
	PrecipitationAmount string `json:"precipitation_amount"`
	WindSpeed           string `json:"wind_speed"`
}

type Meta struct {
	Units     Units     `json:"units"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Timeseries struct {
	Data Data      `json:"data"`
	Time time.Time `json:"time"`
}

type Data struct {
	Instant     Record `json:"instant"`
	Next12Hours Record `json:"next_12_hours"`
	Next1Hours  Record `json:"next_1_hours"`
	Next6Hours  Record `json:"next_6_hours"`
}

type Details struct {
	AirTemperature      float64 `json:"air_temperature"`
	WindSpeed           float64 `json:"wind_speed"`
	PrecipitationAmount float64 `json:"precipitation_amount"`
}

type Summary struct {
	SymbolCode string `json:"symbol_code"`
}

type Record struct {
	Summary Summary `json:"summary,omitempty"`
	Details Details `json:"details"`
}
