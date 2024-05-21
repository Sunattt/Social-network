package handlers

import "github.com/gorilla/mux"

func InitRouter(h *Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/sign-up", h.Registration).Methods("POST")
	router.HandleFunc("/sign-in", h.Login).Methods("POST")

	return router
}
