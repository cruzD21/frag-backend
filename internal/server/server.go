package server

import (
	"fmt"
	"github.com/cruzD21/frag-backend/internal/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	listenAddr string
}

func CreateServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	routes.RegisterFragranceRoutes(router)
	routes.RegisterSearchRoutes(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "server running\n")

	})

	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Printf("http.ListenAndServe -> %v", err)
		return err
	}
	return nil
}
