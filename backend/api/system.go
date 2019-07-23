package api

import (
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
