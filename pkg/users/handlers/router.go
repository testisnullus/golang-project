package handlers

import "github.com/gorilla/mux"

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signUp", h.SignUp())
	r.HandleFunc("/login", h.Login())

	return r
}
