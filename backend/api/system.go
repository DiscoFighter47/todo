package api

import (
	"errors"
	"net/http"

	gson "github.com/DiscoFighter47/gSON"
)

func (api *API) systemCheck(w http.ResponseWriter, r *http.Request) {
	res := gson.Response{
		Data: gson.Object{
			"message": "Hello Universe!",
		},
	}
	res.ServeJSON(w)
}

func (api *API) systemPanic(w http.ResponseWriter, r *http.Request) {
	panic(gson.NewAPIerror("System Panic!", http.StatusInternalServerError, errors.New("system panic")))
}
