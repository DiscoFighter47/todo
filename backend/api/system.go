package api

import "net/http"

func (api *API) systemCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Universe!"))
}
