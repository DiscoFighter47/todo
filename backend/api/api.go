package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// API ...
type API struct {
	handler chi.Router
}

// NewAPI ...
func NewAPI() *API {
	api := &API{}
	api.registerHandler()
	return api
}

// Handler ...
func (api *API) Handler() http.Handler {
	return api.handler
}

func (api *API) registerHandler() {
	api.handler = chi.NewRouter()
	api.handler.Group(func(r chi.Router) {
		r.Mount("/system", api.systemHandlers())
	})
}

func (api *API) systemHandlers() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/check", api.systemCheck)
	})
	return r
}
