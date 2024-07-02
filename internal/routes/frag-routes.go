package routes

import (
	"github.com/cruzD21/frag-backend/internal/controllers"
	"github.com/gorilla/mux"
)

func RegisterFragranceRoutes(router *mux.Router) {
	router.HandleFunc("/frag", controllers.GetFragrances).Methods("GET")
	router.HandleFunc("/frag/{fragID}", controllers.GetFragranceByID).Methods("GET")
}

func RegisterSearchRoutes(router *mux.Router) {
	router.HandleFunc("/search/{query}", controllers.HandleSearch).Methods("GET")
}
