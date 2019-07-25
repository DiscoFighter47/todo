package server

import (
	"fmt"
	"log"
	"net/http"

	gconfig "github.com/DiscoFighter47/gConfig"

	graceful "gopkg.in/tylerb/graceful.v1"

	"github.com/DiscoFighter47/todo/backend/api"
)

// Server ...
type Server struct {
	api    *api.API
	config *gconfig.AppCfg
}

// NewServer ...
func NewServer(api *api.API, config *gconfig.AppCfg) *Server {
	return &Server{
		api:    api,
		config: config,
	}
}

// Serve ...
func (svr *Server) Serve() {
	log.Println("starting server...")

	server := &graceful.Server{
		Timeout: svr.config.GraceTimeout,
		Server: &http.Server{
			ReadTimeout:  svr.config.ReadTimeout,
			WriteTimeout: svr.config.WriteTimeout,
			IdleTimeout:  svr.config.IdelTimeout,
			Addr:         fmt.Sprintf(":%d", svr.config.Port),
			Handler:      svr.api.Handler(),
		},
	}

	log.Println("server listening on port", svr.config.Port)
	log.Fatal(server.ListenAndServe())
}
