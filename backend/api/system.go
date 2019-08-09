package api

import (
	"errors"
	"net/http"

	gson "github.com/DiscoFighter47/gSON"
)

func (api *API) systemCheck(w http.ResponseWriter, r *http.Request) {
	gson.ServeData(w, gson.Object{
		"message": "Hello Universe!",
	})
}

func (api *API) systemPanic(w http.ResponseWriter, r *http.Request) {
	panic(gson.NewAPIerror("System Panic!", http.StatusInternalServerError, errors.New("system test panic")))
}

func (api *API) systemError(w http.ResponseWriter, r *http.Request) {
	a, b := 1, 0
	_ = a / b
}

func (api *API) systemSecret(w http.ResponseWriter, r *http.Request) {
	gson.ServeData(w, gson.Object{
		"message": "Hello Secret Universe! Welcome " + r.Header.Get("subject"),
	})
}
