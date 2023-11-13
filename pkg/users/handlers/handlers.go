package handlers

import (
	"encoding/json"
	"github.com/testisnullus/golang-project/pkg/models"
	"github.com/testisnullus/golang-project/pkg/users/service"
	"io"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandlers(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user := &models.User{}
		err = json.Unmarshal(b, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.services.CreateUser(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		loginUser := &models.LoginUser{}
		err = json.Unmarshal(b, loginUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := h.services.Login(r.Context(), loginUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token)
		w.WriteHeader(http.StatusOK)
	}
}
