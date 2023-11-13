package models

import "time"

type CurrentWeatherResponse struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Wind    Wind
}

type Wind struct {
	Speed float64 `json:"speed"`
}

type Main struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
}

type Weather struct {
	Description string `json:"description"`
}

type CurrentWeather struct {
	ID          int       `json:"ID" db:"id"`
	City        string    `json:"city" db:"city"`
	Description string    `json:"description" db:"description"`
	Time        time.Time `json:"time" db:"time"`
	Speed       float64   `json:"speed" db:"speed"`
	Temp        float64   `json:"temp" db:"temp"`
	Humidity    float64   `json:"humidity" db:"humidity"`
}
