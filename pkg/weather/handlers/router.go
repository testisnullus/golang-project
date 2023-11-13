package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/forecast/now", h.GetWeatherHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/forecast/history", h.GetHistory()).Methods(http.MethodGet)

	return r
}
