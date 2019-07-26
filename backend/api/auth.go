package api

import (
	"net/http"

	gson "github.com/DiscoFighter47/gSON"
	"github.com/DiscoFighter47/todo/backend/model"
)

type authSignUpBody struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (body *authSignUpBody) validate() gson.ValidationError {
	errV := gson.ValidationError{}
	if body.ID == "" {
		errV.Add("id", "required")
	}
	if body.Name == "" {
		errV.Add("name", "required")
	}
	if body.Password == "" {
		errV.Add("password", "required")
	}
	if len(errV) > 0 {
		return errV
	}
	return nil
}

func (api *API) authSignUp(w http.ResponseWriter, r *http.Request) {
	body := &authSignUpBody{}
	if err := gson.ParseBody(r, body); err != nil {
		panic(gson.NewAPIerror("Unable To Parse Body", http.StatusUnprocessableEntity, err))
	}
	if err := body.validate(); err != nil {
		panic(gson.NewAPIerror("Invalid Request Body", http.StatusBadRequest, err))
	}

	if err := api.store.AddUser(&model.User{
		ID:       body.ID,
		Name:     body.Name,
		Password: body.Password,
	}); err != nil {
		panic(gson.NewAPIerror("Unable To Add User", http.StatusBadRequest, err, body.ID))
	}

	res := gson.Response{
		Data: body,
	}
	res.ServeJSON(w)
}
