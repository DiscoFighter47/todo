package server

import (
	"log"
	"net/http"
	"time"

	graceful "gopkg.in/tylerb/graceful.v1"

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

	server := &graceful.Server{
		Timeout: 15 * time.Second,
		Server: &http.Server{
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  15 * time.Second,
			Addr:         ":8080",
			Handler:      svr.api.Handler(),
		},
	}

	log.Println("server listening on port :8080")
	log.Fatal(server.ListenAndServe())
}
