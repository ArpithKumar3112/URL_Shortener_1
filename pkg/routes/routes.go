package routes

import (
	"github.com/arpithku/URL_Shortener_1/pkg/controllers"
	"github.com/gorilla/mux"
)

var SetUpRoutes = func(r *mux.Router) {
	r.HandleFunc("/shorten/{key}/{value}", controllers.PostShortenURL).Methods("POST")
	r.HandleFunc("/shorten/{key}", controllers.GetShortenURL).Methods("GET")
	r.HandleFunc("/shorten/{key}", controllers.DeleteShortenURL).Methods("DELETE")
	r.HandleFunc("/shorten/{key}/{value}", controllers.UpdateShortenURL).Methods("PUT")
}
