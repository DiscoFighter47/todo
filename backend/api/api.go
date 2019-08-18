package api

import (
	"net/http"

	"github.com/DiscoFighter47/todo/backend/cache"

	auth "github.com/DiscoFighter47/gAuth"
	gson "github.com/DiscoFighter47/gSON"
	"github.com/DiscoFighter47/todo/backend/data"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API ...
type API struct {
	handler chi.Router
	store   data.Datastore
	cache   cache.Cache
	auth    *auth.Auth
}

// NewAPI ...
func NewAPI(store data.Datastore, cache cache.Cache, auth *auth.Auth) *API {
	api := &API{
		handler: chi.NewRouter(),
		store:   store,
		cache:   cache,
		auth:    auth,
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
	api.handler.Use(middleware.Logger)
	api.handler.Use(gson.Recoverer)
	api.handler.Use(gson.ReqBodyLogger)
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
		r.Get("/err", api.systemError)
	})
	return r
}

func (api *API) authHandler() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Post("/signup", api.authSignUp)
		r.Post("/signin", api.authSignIn)
		r.With(api.auth.Gatekeeper).Post("/signout", api.authSignOut)
		r.With(api.auth.Gatekeeper).Get("/check", api.authCheck)
	})
	return r
}
