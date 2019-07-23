package server

import (
	"log"
	"net/http"

	"github.com/DiscoFighter47/todo/backend/api"
)

// Server ...
type Server struct {
	api *api.API
}

// NewServer ...
func NewServer(api *api.API) *Server {
	return &Server{
		api: api,
	}
}

// Serve ...
func (svr *Server) Serve() {
	log.Println("starting server...")
	if err := http.ListenAndServe(":8080", svr.api.Handler()); err != nil {
		log.Fatal(err)
	}
}
