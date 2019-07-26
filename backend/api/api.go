package api

import (
	"net/http"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/DiscoFighter47/todo/backend/data"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API ...
type API struct {
	handler chi.Router
	store   data.Datastore
}

// NewAPI ...
func NewAPI(store data.Datastore) *API {
	api := &API{
		handler: chi.NewRouter(),
		store:   store,
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
		r.Mount("/system", api.systemHandler())
		r.Mount("/auth", api.authHandler())
	})
}

func (api *API) systemHandler() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/check", api.systemCheck)
		r.Get("/panic", api.systemPanic)
	})
	return r
}

func (api *API) authHandler() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Post("/signup", api.authSignUp)
	})
	return r
}
