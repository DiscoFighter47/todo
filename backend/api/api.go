package api

import (
	"net/http"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API ...
type API struct {
	handler chi.Router
}

// NewAPI ...
func NewAPI() *API {
	api := &API{
		handler: chi.NewRouter(),
	}
	api.registerMiddleware()
	api.registerHandler()
	return api
}

// Handler ...
func (api *API) Handler() http.Handler {
	return api.handler
}

func (api *API) registerMiddleware() {
	logger := logrus.New()
	api.handler.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	api.handler.Use(gson.Recoverer)
}

func (api *API) registerHandler() {
	api.handler.Group(func(r chi.Router) {
		r.Mount("/system", api.systemHandlers())
	})
}

func (api *API) systemHandlers() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/check", api.systemCheck)
		r.Get("/panic", api.systemPanic)
	})
	return r
}
