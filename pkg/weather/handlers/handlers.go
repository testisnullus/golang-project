package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/testisnullus/golang-project/pkg/config"
	"github.com/testisnullus/golang-project/pkg/jwt"
	"github.com/testisnullus/golang-project/pkg/models"
	"github.com/testisnullus/golang-project/pkg/weather/service"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	cfg      *config.YamlFile
	services *service.Service
}

func NewHandlers(services *service.Service, cfg *config.YamlFile) *Handler {
	return &Handler{services: services, cfg: cfg}
}

func GetCurrentWeather(apiKey, city, date, lat, lon string) (*models.CurrentWeather, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather")

	queryParams := []string{}
	queryParams = append(queryParams, fmt.Sprintf("q=%s", city))

	if len(date) > 0 {
		queryParams = append(queryParams, fmt.Sprintf("dt=%s", date))
	}

	if len(lat) > 0 && len(lon) > 0 {
		queryParams = append(queryParams, fmt.Sprintf("lat=%s", lat))
		queryParams = append(queryParams, fmt.Sprintf("lon=%s", lon))
	}

	queryParams = append(queryParams, fmt.Sprintf("appid=%s", apiKey))
	queryParams = append(queryParams, "units=metric")

	url = fmt.Sprintf("%s?%s", url, strings.Join(queryParams, "&"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	var weatherResponse models.CurrentWeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return nil, err
	}

	currentWeather := &models.CurrentWeather{
		City:        city,
		Description: weatherResponse.Weather[0].Description,
		Time:        time.Now(),
		Speed:       weatherResponse.Wind.Speed,
		Temp:        weatherResponse.Main.Temperature,
		Humidity:    weatherResponse.Main.Humidity,
	}

	return currentWeather, nil
}

func (h *Handler) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get(jwt.AuthHeader)
	_, err := jwt.ParseToken(bearerToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	city := r.URL.Query().Get("city")
	date := r.URL.Query().Get("date")
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	weatherData, err := GetCurrentWeather(h.cfg.ApiKey, city, date, latitude, longitude)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}

	// Extract and return the weather information as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}

func (h *Handler) GetHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get(jwt.AuthHeader)
		_, err := jwt.ParseToken(bearerToken)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		result, err := h.services.Get(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
